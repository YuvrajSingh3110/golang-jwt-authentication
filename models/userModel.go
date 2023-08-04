package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"id"`
	First_name   *string            `json:"first_name" validate:"required,min=2,max=100"`
	Last_name    *string            `json:"last_name" validate:"required,min=2,max=100"`
	Password     *string            `json:"password" validate:"required,min=6"`
	Email        *string            `json:"email" validate:"email,required"`
	Phone        *string            `json:"phone", validate:"required"`
	Token        *string            `json:"token"`
	UserType     *string            `json:"user_type", validate:"required,eq=ADMIN|eq=USER"`
	RefreshToken *string            `json:"refresh_token", validate:"required"`
	CreatedAT    time.Time          `json:"created_at"`
	UpdatedAT    time.Time          `json:"updated_at"`
	User_Id      string             `json:"user_id"`
}
