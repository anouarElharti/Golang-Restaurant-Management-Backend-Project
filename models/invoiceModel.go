package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	ID             primitive.ObjectID `bson:"_id"`
	InvoiceId      string             `json:"invoice_id" validate:"required" bson:"invoice_id"`
	OrderId        string             `json:"order_id" validate:"required" bson:"order_id"`
	Payment_method *string            `json:"payment_method" validate:"required,eq=card|eq=cash" bson:"payment_method"`
	Payment_status *string            `json:"payment_status" validate:"required,eq=paid|eq=unpaid" bson:"payment_status"`
	PaymentDueDate time.Time          `json:"invoice_due_date" validate:"required" bson:"invoice_due_date"`
	Created_at     time.Time          `json:"created_at" validate:"required" bson:"created_at"`
	Updated_at     time.Time          `json:"updated_at" validate:"required" bson:"updated_at"`
}
