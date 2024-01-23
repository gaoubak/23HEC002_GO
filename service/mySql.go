package services

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB

func getDBEnv() (string, string, string, string, string) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var dbUsername = os.Getenv("gonymous")
	var dbPassword = os.Getenv("Goland123")
	var dbHost = os.Getenv("localhost:3306")
	var dbPort = os.Getenv("3306")
	var dbName = os.Getenv("golandDb")

	return dbUsername, dbPassword, dbHost, dbPort, dbName
}

func InitSqlConnection() *gorm.DB {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var dbUsername, dbPassword, dbHost, dbPort, dbName = getDBEnv()

	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	return database
}

func GetConnection() *gorm.DB {
	if database != nil {
		return database
	}

	return InitSqlConnection()
}
