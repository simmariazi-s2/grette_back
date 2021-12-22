package entities

type Board struct {
	Id             int    `gorm:"primaryKey;autoIncrement;index;not null;column:id"`
	Title          string `gorm:"column:title"`
	Contents       string `gorm:"column:contents"`
	CategoryId     int    `gorm:"column:categoryId"`
	CategoryName   string `gorm:"column:categoryName"`
	UserId         string `gorm:"column:userId"`
	CompanyId      int    `gorm:"column:companyId"`
	ReCommandCount int    `gorm:"column:recommandCount"`
	DeCommandCount int    `gorm:"column:decommandCount"`
	DeclareCount   int    `gorm:"column:declareCount"`
}
