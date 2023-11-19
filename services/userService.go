package services

import (
	"context"
	"instant/models"
	"instant/response"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/jeanphorn/log4go"
	"go.mongodb.org/mongo-driver/bson"
)

type Photo struct {
	Photo     []string  `bson:"photo" json:"photo"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

func init() {
	log.LoadConfiguration("./config/logger.json")
}

func UpdateFields(newUser models.User, user models.User) models.User {

	if user.Name == "" {
		user.Name = newUser.Name
	}
	if user.Gender == "" {
		user.Gender = newUser.Gender
	}

	if user.Like_To_Date == "" {
		user.Like_To_Date = newUser.Like_To_Date
	}
	if user.Find_Relation_Type == "" {
		user.Find_Relation_Type = newUser.Find_Relation_Type
	}
	if len(user.Interest) == 0 {
		user.Interest = newUser.Interest
	}
	if user.Smoke == false {
		user.Smoke = newUser.Smoke
	}
	if user.Star_Sign == "" {
		user.Star_Sign = newUser.Star_Sign
	}
	if user.About == "" {
		user.About = newUser.About
	}
	if user.Deactivate_Account == false {
		user.Deactivate_Account = newUser.Deactivate_Account
	}
	if user.Delete_Account == false {
		user.Delete_Account = newUser.Delete_Account
	}

	if user.Country_Code == "" {
		user.Country_Code = newUser.Country_Code
	}

	if user.Email == "" {
		user.Email = newUser.Email
	}

	if user.Role == "" {
		user.Role = "user"
	}

	user.CreatedAt = newUser.CreatedAt

	return user
}

func UserService(c *gin.Context) {

	var user models.User
	var newUser models.User

	// bind req.body to user struct
	if err := c.BindJSON(&user); err != nil {
		log.LOGGER("Test").Error(err.Error())
		response.SendErrorResponse(c, 500, err.Error())
		return
	}

	// checking user is there or not
	count, err := userCollection.CountDocuments(context.TODO(), bson.M{"phone_number": user.Phone_Number})
	if err != nil {
		response.SendErrorResponse(c, 400, err.Error())
		return
	}

	if count == 0 {
		response.SendErrorResponse(c, 400, "user is not found")
		return
	}

	// Find User Data
	if err := userCollection.FindOne(context.TODO(), bson.M{"phone_number": user.Phone_Number}).Decode(&newUser); err != nil {
		response.SendErrorResponse(c, 500, err.Error())
		return
	}

	userData := UpdateFields(newUser, user)

	// Update User
	userData.UpdatedAt = time.Now()
	updateCount, err := userCollection.UpdateOne(context.TODO(), bson.M{"phone_number": user.Phone_Number}, bson.M{"$set": userData})

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

func GetUserByPhoneNumber(c *gin.Context) {
	var user models.User

	if err := userCollection.FindOne(context.TODO(), bson.M{"phone_number": c.Param("phone_number")}).Decode(&user); err != nil {
		response.SendErrorResponse(c, 400, "user is not found")
		return
	}

	response.SendSuccessResponse(c, 200, user)
}

func ViewAllUsers(c *gin.Context) {
	var user []models.User

	cursor, err := userCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		response.SendErrorResponse(c, 400, err.Error())
		return
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var singleProduct models.User
		if err = cursor.Decode(&singleProduct); err != nil {
			response.SendErrorResponse(c, 400, err.Error())
			return
		}
		user = append(user, singleProduct)
	}

	response.SendSuccessResponse(c, 200, user)
}

func UserPhotoService(c *gin.Context) {
	var photo Photo
	phone_number := c.Param("phone_number")

	dst, err := os.Getwd()
	if err != nil {
		response.SendErrorResponse(c, 400, err.Error())
		return
	}

	dst = dst + "/uploads/"

	if _, err := os.Stat(dst); os.IsNotExist(err) {
		println("does not exist")
		if err := os.Mkdir("uploads", os.ModePerm); err != nil {
			response.SendErrorResponse(c, 400, err.Error())
		}
	} else {
		println("The provided directory named")
	}

	form, _ := c.MultipartForm()
	files := form.File["photo[]"]

	var imageUri []string
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	for _, file := range files {

		// Upload the file to specific dst.
		id := uuid.New()
		idString := id.String()
		filePath := idString + filepath.Ext(file.Filename)
		imageUri = append(imageUri, scheme+"://"+c.Request.Host+"/images/"+filePath)
		c.SaveUploadedFile(file, dst+filePath)
	}

	photo.Photo = imageUri
	photo.UpdatedAt = time.Now()

	updateCount, err := userCollection.UpdateOne(context.TODO(), bson.M{"phone_number": phone_number}, bson.M{"$set": photo})

	if err != nil {
		response.SendErrorResponse(c, 400, err.Error())
		return
	}

	if updateCount.ModifiedCount == 0 {
		response.SendErrorResponse(c, 400, "user is not updated")
		return
	}

	response.SendSuccessResponse(c, 200, photo, "message", "user updated successfully")

}
