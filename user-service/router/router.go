package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gtarun8341/ent-go-todo/user-service/ent"
	"github.com/gtarun8341/ent-go-todo/user-service/handlers" // Assuming the correct import path for your handlers
)

func SetupRouter(client ent.Client) *gin.Engine {
	r := gin.Default()

	// Routes for the user service, without the '/users' prefix
	// The API Gateway will handle adding and stripping the '/users' prefix externally.
	r.GET("/status", handlers.Status)
	r.GET("/getAll", handlers.GetAllUsers(&client))
	r.POST("/register", handlers.RegisterUser(&client))
	r.POST("/login", handlers.LoginUser(&client))
	r.GET("/session", handlers.SessionValidate(&client))

	return r
}

