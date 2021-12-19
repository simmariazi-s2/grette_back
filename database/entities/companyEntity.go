package entities

type Company struct {
	Id   int    `gorm:"primaryKey;autoIncrement;index;not null;column:id"`
	Name string `gorm:"column:name"`
}
