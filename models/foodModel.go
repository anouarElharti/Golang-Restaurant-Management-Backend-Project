package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `json:"name" validate:"required,min=2,max=100" bson:"name"`
	Price     *float64           `json:"price" validate:"required" bson:"price"`
	FoodImage string             `json:"foodImage" validate:"required" bson:"foodImage"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	FoodId    string             `json:"foodId" bson:"foodId"`
	MenuId    *string            `json:"menuId" validate:"required" bson:"menuId"`
}
