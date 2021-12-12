package database

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

func ConnectionDB() *gorm.DB {

	err := godotenv.Load("work/grette_back/database/connection.env")

	if err != nil {
		return nil
	}
	dbName := os.Getenv("DB_NAME")
	rootID := os.Getenv("DB_ID")
	rootPW := os.Getenv("DB_PW")
	dbAddr := os.Getenv("DB_ADDR")

	dsn := rootID + ":" + rootPW + "@tcp(" + dbAddr + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"

	gormClient, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	_ = gormClient
	_ = err

	if err != nil {
		return nil
	}

	return gormClient

}
