package repositories

import (
	"work/grette_back/database"
	"work/grette_back/database/entities"
)

// 회사 목록 조회
func GetCompanyList() ([]entities.Company, error) {
	var companyList []entities.Company

	result := database.Db.Model(&companyList).Scan(&companyList)
	/*
		if result != nil {
			return nil, result.Error
		}
	*/

	return companyList, result.Error
}

// 회사 등록
func CreateCompany(company entities.Company) (int, error) {
	result := database.Db.Model(&company).Create(&company)

	return int(result.RowsAffected), result.Error
}
