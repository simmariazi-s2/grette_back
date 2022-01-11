package repositories

import (
	"errors"
	"work/grette_back/database"
	"work/grette_back/database/entities"
)

// 회사 목록 조회
func GetCompanyList() ([]entities.Company, error) {
	var companyList []entities.Company

	database.Db.Model(&companyList).Scan(&companyList)

	if len(companyList) == 0 {
		return nil, errors.New("Company is empty")
	}

	return companyList, nil
}

// 회사 등록
func SetCompany(company entities.Company) (int, error) {
	result := database.Db.Create(&company)

	return int(result.RowsAffected), result.Error
}
