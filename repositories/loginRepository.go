package repositories

import (
	"fmt"
	"time"
	"work/grette_back/database"
	"work/grette_back/database/entities"
)

var UserTable string = "user"
var CompanyTable string = "Company"

type Users struct {
	userId   string
	userName string
}

type Test struct {
	categoryNo int
}

// 이메일로 등록된 유저정보 반환
func GetUser(email string) (entities.User, error) {
	var user entities.User

	//db.Where(&entities.User{Email: email}).First(&user)
	//database.Db.Table("user").Where(&entities.User{Email: email}).Scan(&user)
	database.Db.Model(&user).Where("userId=?", email).Scan(&user)

	return user, nil
}

// 비밀번호 체크
// 동일 비밀번호가 있을 수 있으니, 이메일과 비밀번호로 체크
func ExistsPassword(email string, password string) (int, error) {
	//var user entities.User

	//result := database.Db.Table("user").Where(&entities.User{Email: email, Password: password}).First(&user)

	var existsCount int64 = 1
	//result.Count(&existsCount)

	return int(existsCount), nil
}

func DbTest() string {
	var a string
	var cate []entities.Category
	var user []entities.User
	var b Users
	var d Test
	var iUser entities.User
	//var cmp entities.Company

	//result := database.Db.Table("company")
	//database.Db.Take(&cate)
	//result := database.Db.Table("category").Select("categoryNo", "categoryName").Where("categoryName=?", "테스트").Scan(&cate)
	//database.Db.AutoMigrate()
	//database.Db.First(&cate)
	//if err != nil && err != gorm.ErrRecordNotFound {
	//	return ""
	//}

	//database.Db.Raw("SELECT categoryNo, categoryName FROM category").Scan(&cate)
	//database.Db.Table("company").Scan(&cate)
	database.Db.Model(&entities.Category{}).Limit(10).Find(&d)
	fmt.Println(cate)

	database.Db.Take(&user)
	fmt.Println(user)
	database.Db.Take(&b)
	database.Db.Raw("select * from user").Scan(&user)
	fmt.Println(user)
	fmt.Println(b)
	fmt.Println(d)

	database.Db.Raw("select * from category").Scan(&cate)
	fmt.Println(cate)
	email := "테스트1"
	database.Db.Where("userId=?", email).Find(&user).Scan(&user)
	fmt.Println("@", user)
	database.Db.Model(&user).Where("userId=?", email).Scan(&user)
	fmt.Println("!", user)

	//iUser.UserNo =
	iUser.CompanyNo = 11
	iUser.UserId = "테스트3"
	iUser.UpdateDtm = time.Now()

	//database.Db.Clauses(clause.OnConflict{
	//	Columns:   []clause.Column{{Name: "userId"}},
	//	DoUpdates: clause.AssignmentColumns([]string{"companyNo", "userId", "updateDtm"}),
	//}).Create(&iUser)

	database.Db.Create(&iUser)

	//database.Db.Save(&iUser)
	fmt.Println(iUser.UserNickname)

	//cmp.Domain = "test"
	//cmp.CompanyName = "한글"
	//database.Db.Create(&cmp)

	//fmt.Println("aa :: ", result.Count(&existsCount))

	//a = cate.categoryName + "" + strconv.Itoa(cate.categoryNo)
	//result := database.Db.Table("category").Scan(a)
	// c.uid = "22"
	// c.userName = "33"
	// c.userId = ""
	// c.userNm = ""

	return a
}
