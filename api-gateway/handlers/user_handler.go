package handlers

import "github.com/gin-gonic/gin"

const (
	userServiceBaseURL = "http://localhost:8001"
)


func RegisterUserHandler(c *gin.Context) {
	DynamicProxyHandler(userServiceBaseURL, "/users")(c)
}


func StatusUserHandler(c *gin.Context) {
	DynamicProxyHandler(userServiceBaseURL, "/users")(c)
}


func GetAllUsersHandler(c *gin.Context) {
	DynamicProxyHandler(userServiceBaseURL, "/users")(c)
}


func LoginUserHandler(c *gin.Context) {
	DynamicProxyHandler(userServiceBaseURL, "/users")(c)
}
