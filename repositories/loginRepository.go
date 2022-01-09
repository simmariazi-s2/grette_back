package repositories

import (
	"fmt"
	"strconv"
	"work/grette_back/database"
	"work/grette_back/database/entities"
)

var UserTable string = "user"
var CompanyTable string = "Company"

type Category struct {
	categoryNo   int    `gorm:"primaryKey;autoIncrement;index;not null;column:categoryNo"`
	categoryName string `gorm:"column:categoryName"`
}

type user struct {
	uid      string `gorm:"userId"`
	userName string `gorm:"userName"`
	userId   string `gorm:"default:lemon"`
	userNm   string `gorm:"default:red"`
}

// 이메일로 등록된 유저정보 반환
func GetUser(email string) (entities.User, error) {
	var user entities.User

	//db.Where(&entities.User{Email: email}).First(&user)
	database.Db.Table("user").Where(&entities.User{Email: email}).Scan(&user)

	return user, nil
}

// 비밀번호 체크
// 동일 비밀번호가 있을 수 있으니, 이메일과 비밀번호로 체크
func ExistsPassword(email string, password string) (int, error) {
	var user entities.User

	result := database.Db.Table("user").Where(&entities.User{Email: email, Password: password}).First(&user)

	var existsCount int64
	result.Count(&existsCount)

	return int(existsCount), nil
}

func DbTest() string {
	var a string
	var cate Category
	var c user
	var d user
	//result := database.Db.Table("company")
	//database.Db.Take(&cate)
	//result := database.Db.Table("category").Select("categoryNo", "categoryName").Where("categoryName=?", "테스트").Scan(&cate)
	//database.Db.AutoMigrate()
	//database.Db.First(&cate)
	//if err != nil && err != gorm.ErrRecordNotFound {
	//	return ""
	//}

	database.Db.Raw("SELECT categoryNo, categoryName FROM category").Scan(cate)

	fmt.Println("aa :: ", cate.categoryNo)
	//fmt.Println("aa :: ", result.Count(&existsCount))
	a = cate.categoryName + "" + strconv.Itoa(cate.categoryNo)
	//result := database.Db.Table("category").Scan(a)
	c.uid = "22"
	c.userName = "33"
	c.userId = ""
	c.userNm = ""

	//database.Db.Create(c)
	database.Db.Save(&c)
	//database.Db.Delete(&c)
	gg, _ := database.Setup()
	gg.Create(&c)
	database.Db.Create(&c)
	database.Db.First(d.userId)

	fmt.Println("cc :: ", c)
	fmt.Println("dd :: ", d)
	return a
}
