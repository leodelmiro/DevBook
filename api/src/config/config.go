package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DbConnection = ""
	ApiPort      = 0
)

func LoadEnvironments() {
	var loadEnvironmentsError error

	if loadEnvironmentsError = godotenv.Load(); loadEnvironmentsError != nil {
		log.Fatal(loadEnvironmentsError)
	}

	ApiPort, loadEnvironmentsError = strconv.Atoi(os.Getenv("API_PORT"))
	if loadEnvironmentsError != nil {
		ApiPort = 9000
	}

	DbConnection = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
}
