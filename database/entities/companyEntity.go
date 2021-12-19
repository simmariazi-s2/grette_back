package entities

type Company struct {
	Id   int    `gorm:"column:id"`
	Name string `gorm:"column:name"`
}
