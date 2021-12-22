package setting

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Database struct {
	ConnectionString string
	DbName           string
}

var DatabaseSetting = &Database{}

func Setup() {
	envErr := godotenv.Load(".env")

	if envErr != nil {
		log.Printf(".env file Load Failed")
	}
	dbName := os.Getenv("DB_NAME")
	rootID := os.Getenv("DB_ID")
	rootPW := os.Getenv("DB_PW")
	dbAddr := os.Getenv("DB_ADDR")

	dsn := rootID + ":" + rootPW + "@tcp(" + dbAddr + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"

	DatabaseSetting.ConnectionString = dsn
	DatabaseSetting.DbName = dbName
}
