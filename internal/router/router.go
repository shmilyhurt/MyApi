package router

import (
	"MyApi/internal/api"
	"MyApi/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/register", api.Register)
		apiGroup.POST("/login", api.Login)
		auth := apiGroup.Group("/")
		auth.Use(middleware.JWTAuthMiddleware())
		{
			auth.GET("/users", api.ListUsers)
		}

	}
	return r
}
