package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for GET ALL user"})
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for GET user"})
	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for SignUp"})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for Login"})
	}
}

func HashPassword(password string) string {
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
}
