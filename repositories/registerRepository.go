package repositories

import (
	"errors"
	"work/grette_back/database"
	"work/grette_back/database/entities"

	"gorm.io/gorm"
)

// 이메일 존재 유무 체크
// true: 존재, false: 존재하지 않음
func ExistsUserEmail(email string) (bool, error) {

	var user entities.User
	var result *gorm.DB
	db, err := database.Setup()

	_ = db

	if err != nil {
		return true, err
	}

	// 인자로 이메일값 입력 유무 체크
	if email == "" {
		return true, errors.New("write email is empty")
	}

	// 입력한 이메일에 해당하는 레코드 조회
	result = db.Where(&entities.User{Email: email}).First(&user)

	var existsCount int64
	result.Count(&existsCount)

	// 조회 결과가 0이 아닐경우 true 반환
	if existsCount != 0 {
		return true, nil
	}

	return false, nil
}

// 닉네임 존재 유무 체크
// true: 이미 존재, false: 존재하지 않음
func ExistsNickName(nickName string) (bool, error) {
	var user entities.User
	var result *gorm.DB
	db, err := database.Setup()

	// 데이터베이스 연결 체크
	if err != nil {
		return true, err
	}

	// 닉네임값 입력 체크
	if nickName == "" {
		return true, errors.New("write NickName is empty")
	}

	result = db.Where(&entities.User{NickName: nickName}).First(&user)

	var existsCount int64
	result.Count(&existsCount)

	// 조회 결과가 0이 아닐경우 true 반환
	if existsCount != 0 {
		return true, nil
	}

	return false, nil
}

func GetCompanyList() (map[int]entities.Company, error) {
	var companyList map[int]entities.Company
	db, err := database.Setup()

	if err != nil {
		return nil, err
	}

	db.Table("company").First(companyList)

	if len(companyList) == 0 {
		return nil, errors.New("Company is empty")
	}

	return companyList, nil
}

// 입력된 회원정보 저장
func SetUserRegister(user *entities.User) (bool, error) {
	db, err := database.Setup()

	_ = db

	if err != nil {
		return false, err
	}

	return true, nil
}
