package repositories

import (
	"errors"
	"work/grette_back/database"
	"work/grette_back/database/entities"
	"work/grette_back/util"
)

/*
	이메일로 등록된 유저정보 반환 and 회원가입 여부
	entities.User : 유저 정보 반환
*/
func GetUser(email string) (entities.User, error) {
	var user entities.User

	if email == "" {
		return entities.User{}, errors.New("유저 조회 오류 : 이메일이 비어있습니다.")
	}

	result := database.Db.Model(&user).Where("userId=?", email).Scan(&user)

	return user, result.Error
}

/*
	비밀번호 체크 ::  로그인
	bool : true 로그인 성공, false 로그인 실패
*/
func ExistsPassword(email string, password string) (bool, error) {
	var user entities.User
	var userPassword string

	if email == "" {
		return false, errors.New("로그인 오류 : 아이디(이메일)가 비어있습니다.")
	}

	if password == "" {
		return false, errors.New("로그인 오류 : 비밀번호가 비어있습니다.")
	}

	result := database.Db.Model(&user).Select("userPassword").Where("userId=?", email).Scan(&userPassword)

	return util.CompareHashAndPassword(userPassword, password), result.Error
}

/*
	닉네임 존재 유무 체크
	int : 0 없음, 1 있음
*/
func ExistsNickName(nickName string) (int, error) {
	var user entities.User

	if nickName == "" {
		return 0, errors.New("닉네임 중복체크 오류 : 닉네임이 비어있습니다.")
	}

	result := database.Db.Where("userNickName=?", nickName).Find(&user)

	return int(result.RowsAffected), result.Error
}

// 이메일 존재 유무 체크
// 존재 유무에 다른 개수 반환
func ExistsUserEmail(email string) (int, error) {

	var user entities.User

	// 인자로 이메일값 입력 유무 체크
	if email == "" {
		return 0, errors.New("write email is empty")
	}

	result := database.Db.Where("userId=?", email).Find(&user)

	return int(result.RowsAffected), result.Error
}

/*
	회원가입
	int : 0 실패, 1 성공
*/
func CreateUser(user entities.User) (int, error) {

	if &user == nil {
		return 0, errors.New("회원가입 오류 : User 정보가 비어있습니다. ")
	}

	result := database.Db.Create(&user)

	return int(result.RowsAffected), result.Error
}

/*
	회원정보수정
	int : 0 실패, 1 성공
*/
func UpdateUser(user entities.User) (int, error) {

	if &user == nil {
		return 0, errors.New("회원정보 수정 오류 : User 정보가 비어있습니다. ")
	}

	result := database.Db.Updates(&user)

	return int(result.RowsAffected), result.Error

}
