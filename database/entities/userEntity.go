package entities

import "time"

type User struct {
	UserNo       int        `gorm:"primaryKey;autoIncrement;not null;column:userNo"`
	UserId       *string    `gorm:"not null;column:userId"`
	UserNickname *string    `gorm:"not null;column:userNickname"`
	CompanyNo    *int       `gorm:"not null;column:companyNo"`
	UserPassword *string    `gorm:"not null;column:userPassword"`
	IsUsed       int        `gorm:"not null;column:isUsed;default:1"`
	CreateDtm    *time.Time `gorm:"not null;column:createDtm"`
	UpdateDtm    time.Time  `gorm:"column:updateDtm"`
}
