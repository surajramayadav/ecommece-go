package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id                 primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name               string             `bson:"name" json:"name"`
	Phone_Number       string             `bson:"phone_number" json:"phone_number"  `
	Country_Code       string             `bson:"country_code" json:"country_code"  `
	Email              string             `bson:"email" json:"email"`
	Photo              []string           `bson:"photo" json:"photo"`
	Gender             string             `bson:"gender" json:"gender"`
	Like_To_Date       string             `bson:"like_to_date" json:"like_to_date"`
	Find_Relation_Type string             `bson:"find_relation_type" json:"find_relation_type"`
	Interest           []string           `bson:"interest" json:"interest"`
	Smoke              bool               `bson:"smoke" json:"smoke"`
	Star_Sign          string             `bson:"star_sign" json:"star_sign"`
	About              string             `bson:"about" json:"about"`
	Role               string             `bson:"role" json:"role"`
	Deactivate_Account bool               `bson:"deactivate_account" json:"deactivate_account"`
	Delete_Account     bool               `bson:"delete_account" json:"delete_account"`
	CreatedAt          time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt          time.Time          `bson:"updated_at" json:"updated_at"`
}
