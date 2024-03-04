package routes

import (
	"ridhwankiki/bookstore-go/internals/handlers"
	"ridhwankiki/bookstore-go/internals/middlewares"
	"ridhwankiki/bookstore-go/internals/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitBookRouter(router *gin.Engine, db *sqlx.DB) {
	bookRouter := router.Group("/book")
	bookRepo := repositories.InitBookRepo(db)
	bookHandler := handlers.InitBookHandler(bookRepo)

	// localhost:8000/book
	bookRouter.GET("", bookHandler.GetBooks)
	// localhost:8000/book/new
	bookRouter.POST("/new", middlewares.CheckToken, bookHandler.CreateBook)
	// get book by id
	bookRouter.GET("/:id", middlewares.CheckToken, bookHandler.GetBookById)
	// delete book by id
	bookRouter.DELETE("/:id", middlewares.CheckToken, bookHandler.DeleteBookById)
	// update book by id
	bookRouter.PATCH("/:id", middlewares.CheckToken, bookHandler.UpdateBookById)
	// create book
}
