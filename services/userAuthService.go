package services

import (
	"context"
	"ecommerce/database"
	"ecommerce/models"
	"ecommerce/response"
	"ecommerce/security"
	"ecommerce/utils"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	log "github.com/jeanphorn/log4go"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var Validate = validator.New()

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")

type LoginUserStruct struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func init() {
	log.LoadConfiguration("./config/logger.json")
}

func UserRegistartion(c *gin.Context) {

	var user []models.User
	var singleUser *models.User
	var catchErr string

	singleUser = &models.User{
		Email:    c.PostForm("email"),
		Password: c.PostForm("password"),
		Phone:    c.PostForm("phone"),
		Name:     c.PostForm("name"),
		Role:     c.PostForm("role"),
	}

	if singleUser.Role == "" {
		singleUser.Role = "user"
	}
	// validate user struct
	validatorErr := Validate.Struct(singleUser)
	if validatorErr != nil {
		log.LOGGER("Test").Error(validatorErr.Error())
		response.SendErrorResponse(c, 400, validatorErr.Error())
		return
	}

	// Chceking user already exists or not
	count, err := userCollection.CountDocuments(context.TODO(), bson.M{"email": singleUser.Email})
	if err != nil {
		log.LOGGER("Test").Error(err.Error())
		response.SendErrorResponse(c, 400, err.Error())
		return
	}

	if count > 0 {
		log.LOGGER("Test").Error("email already exists")
		response.SendErrorResponse(c, 400, "Email already exists")
		return
	}

	// convert normal password to hash password
	singleUser.Password, catchErr = utils.HashPassword(singleUser.Password)
	if catchErr != "" {
		log.LOGGER("Test").Error(catchErr)
		response.SendErrorResponse(c, 400, catchErr)
		return
	}

	// multiple Image Upload
	form, _ := c.MultipartForm()
	files := form.File["images"]

	for _, file := range files {
		// Retrieve file Errorrmation
		extension := filepath.Ext(file.Filename)
		// Generate random file name for the new uploaded file so it doesn't override the old file with same name
		newFileName := uuid.New().String() + extension
		singleUser.Photo = append(singleUser.Photo, "http://"+c.Request.Host+"/images/"+newFileName)
		// Upload the file to specific dst.
		if err := c.SaveUploadedFile(file, "./uploads/"+newFileName); err != nil {
			response.SendErrorResponse(c, 500, "Unable to save the file")
			log.LOGGER("Test").Error("Unable to save the file")
			return
		}
	}

	// save user in db
	result, err := userCollection.InsertOne(context.TODO(), singleUser)
	if err != nil {
		log.LOGGER("Test").Error(err.Error())
		response.SendErrorResponse(c, 400, err.Error())
		return
	}

	// checking if user is save or not in db
	if result.InsertedID == "" {
		log.LOGGER("Test").Error("something went wrong")
		response.SendErrorResponse(c, 500, "something went wrong!!!")
		return
	}

	// get user from db
	if err := userCollection.FindOne(context.TODO(), bson.M{"_id": result.InsertedID}).Decode(&singleUser); err != nil {
		response.SendErrorResponse(c, 500, err.Error())
		log.LOGGER("Test").Error(err.Error())
		return
	}

	user = append(user, *singleUser)

	response.SendSuccessResponse(c, 201, user)
}

func UserLogin(c *gin.Context) {
	var user []models.User
	var userLogin LoginUserStruct
	var singleUser models.User

	// bind req.body to user struct
	if err := c.BindJSON(&userLogin); err != nil {
		log.LOGGER("Test").Error(err.Error())
		response.SendErrorResponse(c, 500, err.Error())
		return
	}

	// validate user struct
	validatorErr := Validate.Struct(userLogin)
	if validatorErr != nil {
		log.LOGGER("Test").Error(validatorErr.Error())
		response.SendErrorResponse(c, 400, validatorErr.Error())
		return
	}

	// chceking email or passwprd is null or not
	if userLogin.Email == "" && userLogin.Password == "" {
		log.LOGGER("Test").Error("Email and Password cannot be empty")
		response.SendErrorResponse(c, 400, "Email and Password cannot be empty")
		return
	}

	// get data from db using email
	if err := userCollection.FindOne(context.TODO(), bson.M{"email": userLogin.Email}).Decode(&singleUser); err != nil {
		log.LOGGER("Test").Error("Email and Password is incorrect")
		response.SendErrorResponse(c, 400, "Email and Password is incorrect")
		return
	}

	if singleUser.Email == "" && singleUser.Password == "" {
		log.LOGGER("Test").Error("Email and password is incorrect")
		response.SendErrorResponse(c, 400, "Email and Password is incorrect")
		return
	}

	// convert hash password to normal password getting a flag true of flase
	flag, err := utils.VerifyPassword(singleUser.Password, userLogin.Password)
	if err != "" {
		log.LOGGER("Test").Error("email and password is incorrect")
		response.SendErrorResponse(c, 400, "Email and Password is incorrect")
		return
	}

	if flag == false {
		log.LOGGER("Test").Error("email and password is incorrect")
		response.SendErrorResponse(c, 400, "Email and Password is incorrect")
		return
	}

	// create jwt added email name etc
	jwtToken, err := security.CreateJwtToken(singleUser)

	if err != "" {
		log.LOGGER("Test").Error(err)

		response.SendErrorResponse(c, 500, err)
		return
	}

	user = append(user, singleUser)

	response.SendSuccessResponse(c, 200, user, "token", jwtToken)
}

func UserForgetPassword(c *gin.Context) {
	response.SendErrorResponse(c, 500, "please write forget password code")
}

func UserResetPassword(c *gin.Context) {
	response.SendErrorResponse(c, 500, "please write reset password code")
}

func UserLogout(c *gin.Context) {
	response.SendErrorResponse(c, 500, "please write Logout code")
}
