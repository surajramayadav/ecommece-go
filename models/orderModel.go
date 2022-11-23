package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	Id             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserId         User               `bson:"user" json:"user"`
	Product        Product            `bson:"product" json:"product"`
	Status         string             `bson:"status" json:"status"`
	Payment_status string             `bson:"payment_status" json:"payment_status"`
}
