package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

func ConverIntoObject(id string) (any, string) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {

		return "", "user id is invalid"
	}
	if _id == primitive.NilObjectID {

		return "", "user id is empty"
	}
	return _id, ""

}
