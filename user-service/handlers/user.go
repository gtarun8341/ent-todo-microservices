package handlers

import (
	"net/http"

	"ent-todo-microservices/user-service/ent"
	"ent-todo-microservices/user-service/repositories"

	"github.com/gin-gonic/gin"
)

func Status(c *gin.Context){
	c.String(http.StatusOK, "user service is working")
}
func GetAllUsers(client *ent.Client)gin.HandlerFunc{
	return func(c *gin.Context){
	users,err := repositories.GetAllUsersENT(c.Request.Context(), client)
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"failed to featch users"})
		return
	}
	c.JSON(http.StatusOK,users)
}
}
func SessionValidate(client *ent.Client)gin.HandlerFunc{
	return func(c *gin.Context){
	sessionId := c.Query("token")
	usersId,err := repositories.GetUserIdFromSessionENT(c.Request.Context(), client,sessionId)
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"failed to get session info of users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"userId": usersId})
}
}
func RegisterUser(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input ent.User
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if input.Email == "" || input.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
			return
		}

		err := repositories.RegisterUserENT(c.Request.Context(), client, input )
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"success": "User registered successfully"})
	}
}


func LoginUser(client *ent.Client) gin.HandlerFunc{
	return func(c *gin.Context){
		var input ent.User
			err := c.ShouldBindJSON(&input)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	if input.Email == "" || input.Password == "" {
		c.JSON(http.StatusBadRequest,gin.H{"error":"Email and password are required"})
		return
	}
	session,err :=repositories.LoginUserENT(c.Request.Context(), client, input )
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,session)
	}
}
