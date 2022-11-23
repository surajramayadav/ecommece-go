package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name"`
	Phone    string             `bson:"phone" json:"phone"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
}
