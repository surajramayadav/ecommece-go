package config

import (
	"log"

	"github.com/joho/godotenv"
)

var MONGODB_URI, PORT, GIN_MODE string

func init() {

	var envs map[string]string
	envs, err := godotenv.Read("./.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MONGODB_URI = envs["MONGODB_URI"]
	PORT = envs["PORT"]
	GIN_MODE = envs["GIN_MODE"]

}
