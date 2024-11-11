package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendSingleResponse(ctx *gin.Context, data any, descriptionMsg string, status int) {
	ctx.JSON(http.StatusOK, &SingleResponse{
		Status: Status{
			Code:        status,
			Description: descriptionMsg,
		},
		Data: data,
	})
}

func SendSingleResponseError(ctx *gin.Context, code int, errorMessage string) {
	ctx.AbortWithStatusJSON(code, &Status{
		Code:        code,
		Description: errorMessage,
	})
}
