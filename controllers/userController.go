package controllers

import (
	"context"
	"ecommerce/database"
	"ecommerce/models"
	"ecommerce/response"
	"ecommerce/utils"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gopkg.in/mgo.v2/bson"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		//bind request body to muy struct
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
		count, err := database.Client.Database("ecommerce").Collection("users").CountDocuments(context.TODO(), bson.M{"user": user.Email})
		fmt.Println("count==========", reflect.TypeOf(count), count)
		if err != nil {
			fmt.Println(err)
			response.SendErrorResponse(c, 400, err.Error())
			return
		}

		if count > 0 {
			response.SendErrorResponse(c, 400, "Email already exists")
			return
		}

		user.Password = utils.HashPassword(c, user.Password)
		result, err := database.Client.Database("ecommerce").Collection("users").InsertOne(context.TODO(), user)
		if err != nil {
			response.SendErrorResponse(c, 400, err.Error())
			return
		}
		if result.InsertedID == "" {
			response.SendErrorResponse(c, 500, "something went wrong!!!")
			return
		}
		if err := database.Client.Database("ecommerce").Collection("users").FindOne(context.TODO(), bson.M{"_id": result.InsertedID}).Decode(&user); err != nil {
			response.SendErrorResponse(c, 500, err.Error())
			return
		}
		response.SendSuccessResponse(c, 201, user)

	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		response.SendErrorResponse(c, 401, "login failed")
	}
}

func ViewUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": true, "message": "user view by id successfully"})
	}
}

func ViewUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		response.SendSuccessResponse(c, 200, "user data successfully")
	}
}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": true, "message": "user update successfully"})
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": true, "message": "user delete successfully"})
	}
}
