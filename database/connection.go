package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DbConnector *gorm.DB

func Setup() (*gorm.DB, error) {
	log.Printf("start gorm connection")
	var err error = nil
	envErr := godotenv.Load(".env")

	if envErr != nil {
		log.Printf(".env file Load Failed")
		return nil, envErr
	}

	dbName := os.Getenv("DB_NAME")
	rootID := os.Getenv("DB_ID")
	rootPW := os.Getenv("DB_PW")
	dbAddr := os.Getenv("DB_ADDR")

	dsn := rootID + ":" + rootPW + "@tcp(" + dbAddr + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"

	DbConnector, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("database connection failed")
		return nil, err
	}

	log.Printf("database connection success")

	return DbConnector, err
}
