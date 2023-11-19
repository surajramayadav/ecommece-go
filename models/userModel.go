package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name" validate:"required"`
	Phone    string             `bson:"phone" json:"phone" validate:"required" `
	Email    string             `bson:"email" json:"email" validate:"required"`
	Photo    []string           `bson:"photo" json:"photo"`
	Password string             `bson:"password" json:"password" validate:"required"`
	Role     string             `bson:"role" json:"role"`
}
