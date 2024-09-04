package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO: Implement
	}
}

func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO: Implement
	}
}

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO: Implement
	}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO: Implement
	}
}

func GetOrderItemsByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO: Implement
	}
}

func ItemsByOrder(id string) (OrderItems []primitive.M, err error) {

}
