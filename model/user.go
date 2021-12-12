package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	id    int
	name  string
	email string
}
