package security

import (
	"instant/config"
	"instant/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TokenWithDetails struct {
	Id    primitive.ObjectID
	Name  string
	Phone string
	Email string
	Role  string
	jwt.StandardClaims
}

func CreateJwtToken(user models.User) (string, string) {
	claims := &TokenWithDetails{
		Id:    user.Id,
		Name:  user.Name,
		Phone: user.Phone_Number,
		Email: user.Email,
		Role:  user.Role,
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

func ValidateJwtToken(signedtoken string) (claims *TokenWithDetails, msg string) {

	jwtToken, err := jwt.ParseWithClaims(signedtoken, &TokenWithDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_SECRET), nil
	})

	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := jwtToken.Claims.(*TokenWithDetails)
	if !ok {
		msg = "The Token is invalid"
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "token is expired"
		return
	}
	return claims, msg
}
