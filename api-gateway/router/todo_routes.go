package router

import (
	middleware "ent-todo-microservices/api-gateway/auth"
	"ent-todo-microservices/api-gateway/handlers"

	"github.com/gin-gonic/gin"
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
