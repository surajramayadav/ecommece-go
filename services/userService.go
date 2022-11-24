package services

import (
	"context"
	"ecommerce/database"
	"ecommerce/models"
	"ecommerce/response"
	"ecommerce/security"
	"ecommerce/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var validate = validator.New()

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")

type LoginUserStruct struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func UserRegistartion(c *gin.Context) {

	var user []models.User
	var singleUser *models.User
	var catchErr string

	if err := c.BindJSON(&singleUser); err != nil {
		response.SendErrorResponse(c, 500, err.Error())
		return
	}

	validatorErr := validate.Struct(singleUser)
	if validatorErr != nil {
		response.SendErrorResponse(c, 400, validatorErr.Error())
		return
	}
	count, err := userCollection.CountDocuments(context.TODO(), bson.M{"email": singleUser.Email})
	if err != nil {
		fmt.Println(err)
		response.SendErrorResponse(c, 400, err.Error())
		return
	}

	if count > 0 {
		response.SendErrorResponse(c, 400, "Email already exists")
		return
	}

	singleUser.Password, catchErr = utils.HashPassword(singleUser.Password)
	if catchErr != "" {
		response.SendErrorResponse(c, 400, catchErr)
		return
	}
	result, err := userCollection.InsertOne(context.TODO(), singleUser)
	if err != nil {
		response.SendErrorResponse(c, 400, err.Error())
		return
	}
	if result.InsertedID == "" {
		response.SendErrorResponse(c, 500, "something went wrong!!!")
		return
	}

	if err := userCollection.FindOne(context.TODO(), bson.M{"_id": result.InsertedID}).Decode(&singleUser); err != nil {
		response.SendErrorResponse(c, 500, err.Error())
		return
	}

	user = append(user, *singleUser)

	response.SendSuccessResponse(c, 201, user)
}

func UserLogin(c *gin.Context) {
	var user []models.User
	var userLogin LoginUserStruct
	var singleUser models.User

	if err := c.BindJSON(&userLogin); err != nil {
		response.SendErrorResponse(c, 500, err.Error())
		return
	}
	validatorErr := validate.Struct(userLogin)
	if validatorErr != nil {
		response.SendErrorResponse(c, 400, validatorErr.Error())
		return
	}
	if userLogin.Email == "" && userLogin.Password == "" {
		response.SendErrorResponse(c, 400, "Email and Password cannot be empty")
		return
	}

	if err := userCollection.FindOne(context.TODO(), bson.M{"email": userLogin.Email}).Decode(&singleUser); err != nil {
		fmt.Println(err.Error())
		response.SendErrorResponse(c, 400, "Email and Password is incorrect")
		return
	}
	if singleUser.Email == "" && singleUser.Password == "" {
		response.SendErrorResponse(c, 400, "Email and Password is incorrect")
		return
	}

	flag, err := utils.VerifyPassword(singleUser.Password, userLogin.Password)
	if err != "" {
		response.SendErrorResponse(c, 400, "Email and Password is incorrect")
		return
	}

	if flag == false {
		response.SendErrorResponse(c, 400, "Email and Password is incorrect")
		return
	}

	jwtToken, err := security.CreateJwtToken(singleUser)

	if err != "" {
		response.SendErrorResponse(c, 500, err)
		return
	}

	user = append(user, singleUser)

	response.SendSuccessResponse(c, 200, user, "token", jwtToken)
}

func ViewUserById(c *gin.Context) {

	var user []models.User
	var singleUser models.User

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		response.SendErrorResponse(c, 400, "user id is invalid")
		return
	}
	if id == primitive.NilObjectID {
		response.SendErrorResponse(c, 400, "user id is empty")
		return
	}

	if err := userCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&singleUser); err != nil {
		fmt.Println(err.Error())
		response.SendErrorResponse(c, 400, "user not found")
		return
	}
	user = append(user, singleUser)

	response.SendSuccessResponse(c, 200, user)
}

func ViewUser(c *gin.Context) {

	var user []models.User

	cursor, err := userCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		response.SendErrorResponse(c, 400, err.Error())
		return
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var singleUser models.User
		if err = cursor.Decode(&singleUser); err != nil {
			response.SendErrorResponse(c, 400, err.Error())
			return
		}
		user = append(user, singleUser)
	}

	response.SendSuccessResponse(c, 200, user)

}

func UpdateUser(c *gin.Context) {
	var user []models.User
	var singleUser models.User

	if err := c.BindJSON(&singleUser); err != nil {
		response.SendErrorResponse(c, 500, err.Error())
	}
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		response.SendErrorResponse(c, 400, "user id is invalid")
		return
	}
	if id == primitive.NilObjectID {
		response.SendErrorResponse(c, 400, "user id is empty")
		return
	}

	count, err := userCollection.CountDocuments(context.TODO(), bson.M{"_id": id})
	if err != nil {
		fmt.Println(err)
		response.SendErrorResponse(c, 400, err.Error())
		return
	}

	if count == 0 {
		response.SendErrorResponse(c, 400, "user not found")
		return
	}

	update := bson.M{"name": singleUser.Name, "phone": singleUser.Phone, "email": singleUser.Email}
	updateCount, err := userCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": update})

	if err != nil {
		response.SendErrorResponse(c, 400, err.Error())
		return
	}

	if updateCount.ModifiedCount == 0 {
		response.SendErrorResponse(c, 400, "user is not updated")
		return
	}

	response.SendSuccessResponse(c, 200, user, "message", "user updated successfully")

}

func DeleteUser(c *gin.Context) {
	var user []models.User
	var singleUser models.User

	if err := c.BindJSON(&singleUser); err != nil {
		response.SendErrorResponse(c, 500, err.Error())
	}
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		response.SendErrorResponse(c, 400, "user id is invalid")
		return
	}
	if id == primitive.NilObjectID {
		response.SendErrorResponse(c, 400, "user id is empty")
		return
	}

	count, err := userCollection.CountDocuments(context.TODO(), bson.M{"_id": id})
	if err != nil {
		fmt.Println(err)
		response.SendErrorResponse(c, 400, err.Error())
		return
	}

	if count == 0 {
		response.SendErrorResponse(c, 400, "user not found")
		return
	}
	deleteCount, err := userCollection.DeleteOne(context.TODO(), bson.M{"_id": id})

	if err != nil {
		response.SendErrorResponse(c, 400, err.Error())
		return
	}

	if deleteCount.DeletedCount == 0 {
		response.SendErrorResponse(c, 400, "user is not deleted")
		return
	}

	response.SendSuccessResponse(c, 200, user, "message", "user deleted successfully")

}
