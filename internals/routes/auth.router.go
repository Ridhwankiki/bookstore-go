package routes

import (
	"ridhwankiki/bookstore-go/internals/handlers"
	"ridhwankiki/bookstore-go/internals/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitAuthRouter(router *gin.Engine, db *sqlx.DB) {
	// Bikin Sub-router
	authRouter := router.Group("/auth")
	authRepo := repositories.InitAuthRepo(db)
	authHandler := handlers.InitAuthHandler(authRepo)

	// Bikin Rute nya
	// Localhost:8000/auth/new
	authRouter.POST("/new", authHandler.Register)
	// Localhost:8000/auth
	authRouter.POST("", authHandler.Login)
}
