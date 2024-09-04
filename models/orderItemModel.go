package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Quantity    *string            `bson:"quantity" validate="required,eq=S|eq=M|eq=L" json:"quantity"`
	UnitPrice   *float64           `bson:"unit_price" validate="required" json:"unit_price"`
	OrderItemId string             `bson:"order_item_id" json:"order_item_id"`
	OrderId     string             `bson:"order_id" json:"order_id"`
	FoodId      *string            `bson:"food_id" validate="required" json:"food_id"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}
