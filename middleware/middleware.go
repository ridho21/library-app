package middleware

import (
	"net/http"
	"strings"
	"test-ordent/model/dto/response"
	"test-ordent/utils/common"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware interface {
	RequireToken(roles ...string) gin.HandlerFunc
}

type authMiddleware struct {
	jwtToken common.JwtToken
}

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

func (a *authMiddleware) RequireToken(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var aH authHeader
		if err := c.ShouldBindHeader(&aH); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Status{
				Code:        http.StatusUnauthorized,
				Description: err.Error(),
			})
			return
		}

		tokenString := strings.Replace(aH.AuthorizationHeader, "Bearer ", "", -1)
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Status{
				Code:        http.StatusUnauthorized,
				Description: "Unauthorized",
			})
			return
		}
		claims, err := a.jwtToken.VerifyToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Status{
				Code:        http.StatusUnauthorized,
				Description: err.Error(),
			})
			return
		}

		validRole := false
		if len(roles) > 0 {
			for _, role := range roles {
				if role == claims["role"] {
					validRole = true
					break
				}
			}
		}

		if !validRole {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Status{
				Code:        http.StatusUnauthorized,
				Description: "You don't have permission",
			})
			return
		}
		c.Next()
	}
}

func NewAuthMiddleware(jwtToken common.JwtToken) AuthMiddleware {
	return &authMiddleware{
		jwtToken: jwtToken,
	}
}
