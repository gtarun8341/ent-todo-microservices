package router

import (
	"ent-todo-microservices/api-gateway/handlers"

	"github.com/gin-gonic/gin"
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
