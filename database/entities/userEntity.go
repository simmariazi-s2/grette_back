package entities

type User struct {
	Id          int    `gorm:"primaryKey;autoIncrement;index;not null;column:id"`
	NickName    string `gorm:"column:nickName"`
	Email       string `gorm:"column:email"`
	Password    string `gorm:"column:password"`
	CompanyName string `gorm:"column:companyName"`
}
