package services

import (
	"context"
	"instant/database"
	"instant/models"
	"instant/response"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	log "github.com/jeanphorn/log4go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var Validate = validator.New()

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")

func init() {
	log.LoadConfiguration("./config/logger.json")
}

func AuthService(c *gin.Context) {

	var user models.User
	// bind req.body to user struct
	if err := c.BindJSON(&user); err != nil {
		log.LOGGER("Test").Error(err.Error())
		response.SendErrorResponse(c, 500, err.Error())
		return
	}

	// checking dublication of field
	count, err := userCollection.CountDocuments(context.TODO(), bson.M{"phone_number": user.Phone_Number})
	if err != nil {
		response.SendErrorResponse(c, 400, err.Error())
		return
	}

	if count > 0 {
		response.SendErrorResponse(c, 400, "phone number already exists")
		return
	}

	// Save PhoneNumber in db
	user.CreatedAt = time.Now()
	user.Role = "user"
	result, err := userCollection.InsertOne(context.TODO(), user)

	if err != nil {
		response.SendErrorResponse(c, 400, err.Error())
		return
	}

	if result.InsertedID == "" {
		response.SendErrorResponse(c, 400, err.Error())
		return
	}

	// Find User Data
	if err := userCollection.FindOne(context.TODO(), bson.M{"_id": result.InsertedID}).Decode(&user); err != nil {
		response.SendErrorResponse(c, 500, err.Error())
		return
	}

	response.SendSuccessResponse(c, 200, user)

}
