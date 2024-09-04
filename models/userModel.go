package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	FirstName    *string            `json:"name" validate="required,min=2,max=100" bson:"name"`
	LastName     *string            `json:"lastName" validate:"required,min=2,max=100" bson:"lastName"`
	Email        *string            `json:"email" validate:"email,required" bson:"email"`
	AvatarUrl    *string            `json:"avatarUrl" bson:"avatarUrl"`
	Phone        *string            `json:"phone" validate:"required" bson:"phone"`
	Password     *string            `json:"Password" validate="required,min=8" bson:"password"`
	Token        *string            `json:"token" bson:"token"`
	RefreshToken *string            `json:"refreshToken" bson:"refreshToken"`
	CreatedAt    time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt" bson:"updatedAt"`
	UserId       string             `json:"user_id" bson:"userId"`
}
