package database

import (
	"log"
	"work/grette_back/database/entities"
	"work/grette_back/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var Db *gorm.DB

// func Setup() (*gorm.DB, error) {
// 	var err error = nil

// 	log.Printf("데이터베이스 시작!")
// 	Db, err = gorm.Open(mysql.Open(setting.DatabaseSetting.ConnectionString), &gorm.Config{})
// 	if err != nil {
// 		log.Printf("데이터베이스 실패 %s", err.Error())
// 		return nil, err
// 	}

// 	err = Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&entities.User{})

// 	if err != nil {
// 		log.Printf("데이터베이스 실패 : %s", err.Error())
// 		return nil, err
// 	}
// 	log.Printf("데이터베이스 연결 성공")
// 	return Db, nil
// }

var Db *gorm.DB

func Setup() (*gorm.DB, error) {

	var err error = nil

	log.Printf("start gorm connection")
	Db, err := gorm.Open(mysql.Open(setting.DatabaseSetting.ConnectionString), &gorm.Config{})

	if err != nil {
		log.Printf("database connection failed: %s", err.Error())
		return nil, err
	}

	err = Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&entities.User{})

	if err != nil {
		log.Printf("database connection failed: %s", err.Error())
		return nil, err
	}

	log.Printf("database connection success")

	return Db, nil
}
