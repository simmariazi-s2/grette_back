package main

import (
	"work/grette_back/database"
)

type User struct {
	Id          int    `gorm:"primaryKey;autoIncrement;index;not null;column:id"`
	Name        string `gorm:"column:name"`
	Email       string `gorm:"column:email"`
	Password    string `gorm:"column:password"`
	CompanyName string `gorm:"column:companyName"`
}

func main() {

	db, err := database.Setup()

	var user User
	aa := db.Select("email").Where(User{Email: "baekhk1006@gmail.com"}).First(&user)

	_ = aa
	_ = db
	_ = err
}
