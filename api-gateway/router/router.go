package router

import (
	"github.com/gin-gonic/gin"
)


func RootStatusHandler(c *gin.Context) {
	c.JSON(200, gin.H{"status": "API Gateway running"})
}

func SetupRouter() *gin.Engine {
	r := gin.Default()


	r.GET("/", RootStatusHandler) 

	RegisterTodoRoutes(r)
	RegisterUserRoutes(r)

	return r
}
