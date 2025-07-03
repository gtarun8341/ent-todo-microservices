package router

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/gtarun8341/ent-todo-microservices/api-gateway/auth"
	"github.com/gtarun8341/ent-todo-microservices/api-gateway/handlers"
)

func RegisterTodoRoutes(r *gin.Engine) {
	todoGroup := r.Group("/todo")
	todoGroup.Use(middleware.AuthMiddleware())
	{
		todoGroup.GET("/status", handlers.StatusTodoHandler)
		todoGroup.GET("/", handlers.GetAllTodosHandler)
		todoGroup.POST("/", handlers.CreateTodoHandler)
		todoGroup.GET("/:id", handlers.GetTodoByIdHandler)
		todoGroup.PUT("/:id", handlers.UpdateTodoByIdHandler)
		todoGroup.DELETE("/:id", handlers.DeleteTodoByIdHandler)
	}
}
