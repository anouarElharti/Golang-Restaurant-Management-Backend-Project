package controllers

import (
	"context"
	"golang-restaurant-management/database"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionMenu *mongo.Collection = database.OpenCollection(database.Client, "menu")

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := collectionMenu.Find(context.TODO(), bson.M{})
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while fetching the menus!"})
		}

		var allMenus []bson.M
		if err = result.All(ctx, &allMenus); err != nil {
			log.Fatal("Error : Error occured while fetching the menus!")
		}

		c.JSON(http.StatusOK, allMenus)

	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO: Implement
	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO: Implement
	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO: Implement
	}
}
