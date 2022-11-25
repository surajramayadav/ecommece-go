package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	Id             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	User           string             `bson:"user" json:"user"`
	Product        string             `bson:"product" json:"product"`
	Status         bool               `bson:"status" json:"status"`
	Payment_status bool               `bson:"payment_status" json:"payment_status"`
}
