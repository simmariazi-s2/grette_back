package entities

import "time"

type Company struct {
	CompanyNo   int       `gorm:"primaryKey;autoIncrement;not null;column:companyNo"`
	CompanyName string    `gorm:"not null;column:companyName"`
	Domain      string    `gorm:"not null;column:domain"`
	Description string    `gorm:"not null;column:description"`
	IsUsed      int       `gorm:"not null;column:isUsed;default:1"`
	CreateDtm   time.Time `gorm:"not null;column:createDtm"`
}
