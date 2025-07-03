package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gtarun8341/ent-todo-microservices/api-gateway/handlers"
)

func RegisterUserRoutes(r *gin.Engine) {
	userGroup := r.Group("/users")
	{
		userGroup.GET("/status", handlers.StatusUserHandler)
		userGroup.GET("/getAll", handlers.GetAllUsersHandler)
		userGroup.POST("/register", handlers.RegisterUserHandler)
		userGroup.POST("/login", handlers.LoginUserHandler)
	}
}
