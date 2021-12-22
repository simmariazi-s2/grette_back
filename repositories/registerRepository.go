package repositories

import (
	"errors"
	"work/grette_back/database"
	"work/grette_back/database/entities"
)

// 이메일 존재 유무 체크
// 존재 유무에 다른 개수 반환
func ExistsUserEmail(email string) (int, error) {

	var user entities.User

	// 인자로 이메일값 입력 유무 체크
	if email == "" {
		return 0, errors.New("write email is empty")
	}

	// 입력한 이메일에 해당하는 레코드 조회
	//result = database.DbConnector.Where(&entities.User{Email: email}).First(&user)
	result := database.DbConnector.Table(UserTable).Where(&entities.User{Email: email}).First(&user)

	var existsCount int64
	result.Count(&existsCount)

	// 조회 결과가 0이 아닐경우 true 반환
	if existsCount != 0 {
		return int(existsCount), nil
	}

	return 0, nil
}

// 닉네임 존재 유무 체크
// 존재 유무에 다른 개수 반환
func ExistsNickName(nickName string) (int, error) {
	var user entities.User

	result := database.DbConnector.Table(UserTable).Where(&entities.User{NickName: nickName}).Scan(&user)

	var existsCount int64
	result.Count(&existsCount)

	return int(existsCount), nil
}

// 회사 정보리스트 조회
func GetCompanyList() (map[int]entities.Company, error) {
	var companyList map[int]entities.Company

	database.DbConnector.Table(CompanyTable).Scan(&companyList)

	if len(companyList) == 0 {
		return nil, errors.New("Company is empty")
	}

	return companyList, nil
}

// 입력된 회원정보 저장
func SetUserRegister(user entities.User) (bool, error) {

	result := database.DbConnector.Create(&user)

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
