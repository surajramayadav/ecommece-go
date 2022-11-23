package utils

import (
	"ecommerce/response"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(c *gin.Context, password string) string {
	byte, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		response.SendErrorResponse(c, 500, err.Error())
	}
	return string(byte)
}

func VerifyPassword(c *gin.Context, givenPassword string, userPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(givenPassword), []byte(userPassword))
	valid := true
	if err != nil {
		valid = false
		response.SendErrorResponse(c, 500, err.Error())
	}
	return valid

}
