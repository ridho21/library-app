package controller

import (
	"net/http"
	"test-ordent/middleware"
	"test-ordent/model"
	"test-ordent/model/dto/response"
	"test-ordent/usecase"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	uc             usecase.BookUsecase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (c *BookController) getAllBookHandler(ctx *gin.Context) {
	books, err := c.uc.GetAllBooks()
	if err != nil {
		response.SendSingleResponseError(
			ctx, http.StatusBadRequest, err.Error(),
		)
		return
	}

	var data []any

	data = append(data, books)

	response.SendSingleResponse(
		ctx,
		data,
		"Success",
		http.StatusOK)
}

func (c *BookController) insertNewBook(ctx *gin.Context) {
	var newBook model.Books

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	book, err := c.uc.InsertNewBook(newBook)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Success Insert New Book",
		"data":    book,
	})
}

func (c *BookController) updateBook(ctx *gin.Context) {
	var book model.Books

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := c.uc.UpdateBooks(book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Success Update Book",
		"data":    "",
	})
}

func (c *BookController) deleteBook(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.uc.DeleteBooks(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Success Delete Book",
		"data":    "",
	})
}

func (uc *BookController) Route() {
	router := uc.rg.Group("/book")
	router.Use(uc.authMiddleware.RequireToken("ADMIN"))
	router.GET("", uc.getAllBookHandler)
	router.POST("", uc.insertNewBook)
	router.PUT("", uc.updateBook)
	router.DELETE("/:id", uc.deleteBook)
}

func NewBookController(uc usecase.BookUsecase, router *gin.Engine, authMiddleware middleware.AuthMiddleware) *BookController {
	return &BookController{
		uc:             uc,
		rg:             &router.RouterGroup,
		authMiddleware: authMiddleware,
	}
}
