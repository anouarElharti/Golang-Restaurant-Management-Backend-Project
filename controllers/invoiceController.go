package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for GET ALL invoice"})
	}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for GET invoice"})
	}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for Create invoice"})
	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for Update invoice"})
	}
}
