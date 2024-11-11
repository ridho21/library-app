package controller

import (
	"net/http"
	"test-ordent/model"
	"test-ordent/model/dto/request"
	"test-ordent/model/dto/response"
	"test-ordent/usecase"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	uc usecase.UserUsecase
	rg *gin.RouterGroup
}

func (uc *UserController) registerHandler(ctx *gin.Context) {
	var newUser model.User
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		response.SendSingleResponseError(
			ctx,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}
	user, err := uc.uc.CreateUser(newUser)
	if err != nil {
		response.SendSingleResponseError(
			ctx,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}
	response.SendSingleResponse(
		ctx,
		user,
		"Success Register new User",
		http.StatusCreated,
	)
}

func (uc *UserController) loginHandler(ctx *gin.Context) {
	var loginPayload request.LoginRequestDto
	if err := ctx.ShouldBindJSON(&loginPayload); err != nil {
		response.SendSingleResponseError(
			ctx,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}
	user, err := uc.uc.LoginUser(loginPayload)
	if err != nil {
		response.SendSingleResponseError(
			ctx,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}
	response.SendSingleResponse(
		ctx,
		user,
		"Success Login as User",
		http.StatusCreated,
	)
}

func (uc *UserController) Route() {
	router := uc.rg.Group("/auth")
	router.POST("/register", uc.registerHandler)
	router.POST("/login", uc.loginHandler)
}

func NewUserController(uc usecase.UserUsecase, router *gin.Engine) *UserController {
	return &UserController{
		uc: uc,
		rg: &router.RouterGroup,
	}
}
