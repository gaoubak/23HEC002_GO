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

	var dbUsername = os.Getenv("DB_USERNAME")
	var dbPassword = os.Getenv("DB_PASSWORD")
	var dbHost = os.Getenv("DB_HOST")
	var dbPort = os.Getenv("DB_PORT")
	var dbName = os.Getenv("DB_NAME")

	return dbUsername, dbPassword, dbHost, dbPort, dbName
}

func InitSqlConnection() (*gorm.DB, error) {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	var dbUsername, dbPassword, dbHost, dbPort, dbName = getDBEnv()

	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
		return nil, err
	}

	return database, nil
}

func GetConnection() *gorm.DB {
	if database != nil {
		return database
	}

	_, _ = InitSqlConnection()

	return database
}
