package services

import (
	"context"
	"ecommerce/models"
	"ecommerce/response"
	"ecommerce/utils"

	"github.com/gin-gonic/gin"
	log "github.com/jeanphorn/log4go"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	log.LoadConfiguration("./config/logger.json")
}

func ViewUserById(c *gin.Context) {

	var user []models.User
	var singleUser models.User

	id, err := utils.ConverIntoObject(c.Param("id"))
	if err != "" {
		log.LOGGER("Test").Error(err)
		response.SendErrorResponse(c, 400, err)
		return
	}

	// get data from db
	if err := userCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&singleUser); err != nil {
		log.LOGGER("Test").Error(err.Error())
		response.SendErrorResponse(c, 400, "user is not found")
		return
	}
	user = append(user, singleUser)

	response.SendSuccessResponse(c, 200, user)
}

func ViewUser(c *gin.Context) {

	var user []models.User

	cursor, err := userCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.LOGGER("Test").Error(err.Error())
		response.SendErrorResponse(c, 400, err.Error())
		return
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var singleUser models.User
		if err = cursor.Decode(&singleUser); err != nil {
			log.LOGGER("Test").Error(err.Error())
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
		log.LOGGER("Test").Error(err.Error())
		response.SendErrorResponse(c, 500, err.Error())
	}

	id, e := utils.ConverIntoObject(c.Param("id"))
	if e != "" {
		log.LOGGER("Test").Error(e)
		response.SendErrorResponse(c, 400, e)
		return
	}

	count, err := userCollection.CountDocuments(context.TODO(), bson.M{"_id": id})
	if err != nil {
		log.LOGGER("Test").Error(err.Error())
		response.SendErrorResponse(c, 400, err.Error())
		return
	}

	if count == 0 {
		log.LOGGER("Test").Error("user is not found")
		response.SendErrorResponse(c, 400, "user is not found")
		return
	}

	update := bson.M{"name": singleUser.Name, "phone": singleUser.Phone, "email": singleUser.Email}
	updateCount, err := userCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": update})

	if err != nil {
		response.SendErrorResponse(c, 400, err.Error())
		log.LOGGER("Test").Error(err.Error())
		return
	}

	if updateCount.ModifiedCount == 0 {
		log.LOGGER("Test").Error("user is not updated")
		response.SendErrorResponse(c, 400, "user is not updated")
		return
	}

	response.SendSuccessResponse(c, 200, user, "message", "user updated successfully")

}

func DeleteUser(c *gin.Context) {
	var user []models.User

	id, e := utils.ConverIntoObject(c.Param("id"))
	if e != "" {
		response.SendErrorResponse(c, 400, e)
		log.LOGGER("Test").Error(e)

		return
	}

	count, err := userCollection.CountDocuments(context.TODO(), bson.M{"_id": id})
	if err != nil {
		log.LOGGER("Test").Error(err.Error())
		response.SendErrorResponse(c, 400, err.Error())
		return
	}

	if count == 0 {
		log.LOGGER("Test").Error("user is not found")
		response.SendErrorResponse(c, 400, "user is not found")
		return
	}
	deleteCount, err := userCollection.DeleteOne(context.TODO(), bson.M{"_id": id})

	if err != nil {
		log.LOGGER("Test").Error(err.Error())
		response.SendErrorResponse(c, 400, err.Error())
		return
	}

	if deleteCount.DeletedCount == 0 {
		log.LOGGER("Test").Error("user is not deleted")
		response.SendErrorResponse(c, 400, "user is not deleted")
		return
	}

	response.SendSuccessResponse(c, 200, user, "message", "user deleted successfully")

}
