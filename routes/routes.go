package routes

import (
	"github.com/feribeirods/voz-api/handlers"
	"github.com/feribeirods/voz-api/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Agrupa rotas com prefixo /api/v1
	v1 := r.Group("/api/v1")

	// Rotas públicas (sem autenticação)
	v1.POST("/register", handlers.RegisterHandler)
	v1.POST("/login", handlers.LoginHandler)

	// Rotas protegidas (requer token JWT)
	auth := v1.Group("/")
	auth.Use(middlewares.AuthMiddleware())
	{
		auth.GET("/me", handlers.MeHandler)
	}
}
