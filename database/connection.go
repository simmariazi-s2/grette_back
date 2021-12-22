package database

import (
	"log"
	"work/grette_back/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Setup() (*gorm.DB, error) {
	log.Printf("start gorm connection")
	var err error = nil

	Db, err = gorm.Open(mysql.Open(setting.DatabaseSetting.ConnectionString), &gorm.Config{})

	if err != nil {
		log.Printf("database connection failed")
		return nil, err
	}

	log.Printf("database connection success")

	return Db, err
}
