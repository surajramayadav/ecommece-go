package security

import (
	"ecommerce/config"
	"ecommerce/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TokenWithDetails struct {
	Id    primitive.ObjectID
	Name  string
	Phone string
	Email string
	jwt.StandardClaims
}

func CreateJwtToken(user models.User) (string, string) {
	claims := &TokenWithDetails{
		Id:    user.Id,
		Name:  user.Name,
		Phone: user.Phone,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.JWT_SECRET))

	if err != nil {
		return "", err.Error()
	}

	return token, ""
}
