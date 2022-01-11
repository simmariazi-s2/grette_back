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

	r := database.Db.Where("userId=?", "테스트1").Find(&user)
	fmt.Println("이야호 :: ", r.RowsAffected)
	database.Db.Raw("select * from category").Scan(&cate)
	fmt.Println(cate, " ::::::  ", len(cate))
	email := "테스트1"
	database.Db.Where("userId=?", email).Find(&user).Scan(&user)
	fmt.Println("@", user)

	database.Db.Model(&user).Where("userId=?", email).Update("userPassword", "비밀번호")
	database.Db.Model(&user).Where("userId=?", email).Scan(&user)
	result := map[string]interface{}{}

	database.Db.Model(&user).Select("userPassword").Where("userId=?", email).Scan(&a)
	fmt.Println("! 2", user)
	fmt.Println("! 3", result)
	fmt.Println("! 4", a)

	//iUser.UserNo =
	iUser.CompanyNo = 11
	iUser.UserId = "테스트3"
	iUser.UpdateDtm = time.Now()

	//database.Db.Clauses(clause.OnConflict{
	//	Columns:   []clause.Column{{Name: "userId"}},
	//	DoUpdates: clause.AssignmentColumns([]string{"companyNo", "userId", "updateDtm"}),
	//}).Create(&iUser)

	//database.Db.Create(&iUser)

	//database.Db.Save(&iUser)
	//fmt.Println(iUser.UserNickname)

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

	database.Db.Model(&user).Scan(&user)

	fmt.Println("유저 정보 리스트 :: ", user[0])
	fmt.Println("유저 정보 리스트 :: ", user[1])

	return a
}
