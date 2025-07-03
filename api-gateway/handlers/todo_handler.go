package handlers

import (
	"github.com/gin-gonic/gin"
)
const (
	todoServiceBaseURL = "http://localhost:8002"
)

func StatusTodoHandler(c *gin.Context) {
	DynamicProxyHandler(todoServiceBaseURL, "/todo")(c)
}


func GetAllTodosHandler(c *gin.Context) {
	DynamicProxyHandler(todoServiceBaseURL, "/todo")(c)
}


func CreateTodoHandler(c *gin.Context) {
	DynamicProxyHandler(todoServiceBaseURL, "/todo")(c)
}

func GetTodoByIdHandler(c *gin.Context) {
	DynamicProxyHandler(todoServiceBaseURL, "/todo")(c)
}

func UpdateTodoByIdHandler(c *gin.Context) {
	DynamicProxyHandler(todoServiceBaseURL, "/todo")(c)
}


func DeleteTodoByIdHandler(c *gin.Context) {
	DynamicProxyHandler(todoServiceBaseURL, "/todo")(c)
}