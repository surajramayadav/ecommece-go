package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, string) {
	byte, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return string(byte), err.Error()
	}
	return string(byte), ""
}

func VerifyPassword(userPassword string, givenPassword string) (bool, string) {
	// first is hash password and second parameter is password

	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(givenPassword))

	if err != nil {
		return false, err.Error()
	}
	return true, ""

}
