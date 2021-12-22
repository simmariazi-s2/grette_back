package repositories

import (
	"work/grette_back/database"
	"work/grette_back/database/entities"
)

var UserTable string = "user"
var CompanyTable string = "Company"

// 이메일로 등록된 유저정보 반환
func GetUser(email string) (entities.User, error) {
	var user entities.User

	//db.Where(&entities.User{Email: email}).First(&user)
	database.Db.Table(UserTable).Where(&entities.User{Email: email}).Scan(&user)

	return user, nil
}

// 비밀번호 체크
// 동일 비밀번호가 있을 수 있으니, 이메일과 비밀번호로 체크
func ExistsPassword(email string, password string) (int, error) {
	var user entities.User

	result := database.Db.Table(UserTable).Where(&entities.User{Email: email, Password: password}).First(&user)

	var existsCount int64
	result.Count(&existsCount)

	return int(existsCount), nil
}
