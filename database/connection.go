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

	//db, err := gorm.Open(mysql.Open(setting.DatabaseSetting.ConnectionString), &gorm.Config{})
	//	db, err := gorm.Open("mysql", mysql.Open(setting.DatabaseSetting.ConnectionString))

	Db, err = gorm.Open(mysql.Open(setting.DatabaseSetting.ConnectionString), &gorm.Config{})

	if err != nil {
		log.Printf("database connection failed err :: ", err)
		return nil, err
	}
	/*
	 DBCP(DataBaseConnectionPool) 설정
	 SetMaxIdleConns : 대기 커넥션
	 SetMaxOpenConns : 최대 커넥션
	 SetConnMaxLifetime : 커넥션 최대 유지시간
	*/
	//sqlDb, _ := Db.DB()

	//sqlDb.SetMaxIdleConns(10)
	//sqlDb.SetMaxOpenConns(100)
	//sqlDb.SetConnMaxLifetime(time.Hour)

	log.Printf("database connection success")

	return Db, err
}
