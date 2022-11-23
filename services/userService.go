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
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var validate = validator.New()

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")

type LoginUserStruct struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	User  models.User
	Token string `json:"token"`
}

func UserRegistartion(c *gin.Context) {

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		response.SendErrorResponse(c, 400, err.Error())
		return

	}
	validatorErr := validate.Struct(user)
	if validatorErr != nil {
		response.SendErrorResponse(c, 400, validatorErr.Error())
		return
	}
	count, err := userCollection.CountDocuments(context.TODO(), bson.M{"email": user.Email})
	if err != nil {
		fmt.Println(err)
		response.SendErrorResponse(c, 400, err.Error())
		return
	}

	if count > 0 {
		response.SendErrorResponse(c, 400, "Email already exists")
		return
	}

	var catchErr string
	user.Password, catchErr = utils.HashPassword(user.Password)
	if catchErr != "" {
		response.SendErrorResponse(c, 400, catchErr)
		return
	}
	result, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		response.SendErrorResponse(c, 400, err.Error())
		return
	}
	if result.InsertedID == "" {
		response.SendErrorResponse(c, 500, "something went wrong!!!")
		return
	}
	if err := userCollection.FindOne(context.TODO(), bson.M{"_id": result.InsertedID}).Decode(&user); err != nil {
		response.SendErrorResponse(c, 500, err.Error())
		return
	}
	response.SendSuccessResponse(c, 201, user)
}

func UserLogin(c *gin.Context) {
	var userLogin LoginUserStruct
	if err := c.BindJSON(&userLogin); err != nil {
		response.SendErrorResponse(c, 400, err.Error())
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
	var user models.User
	if err := userCollection.FindOne(context.TODO(), bson.M{"email": userLogin.Email}).Decode(&user); err != nil {
		fmt.Println(err.Error())
		response.SendErrorResponse(c, 400, "Email and Password is incorrect")
		return
	}
	if user.Email == "" && user.Password == "" {
		response.SendErrorResponse(c, 400, "Email and Password is incorrect")
		return
	}

	flag, err := utils.VerifyPassword(user.Password, userLogin.Password)
	if err != "" {
		response.SendErrorResponse(c, 400, "Email and Password is incorrect")
		return
	}

	if flag == false {
		response.SendErrorResponse(c, 400, "Email and Password is incorrect")
		return
	}

	jwtToken, err := security.CreateJwtToken(user)

	if err != "" {
		response.SendErrorResponse(c, 500, err)
		return
	}

	res := LoginResponse{
		User:  user,
		Token: jwtToken,
	}

	response.SendSuccessResponse(c, 200, res)
}
