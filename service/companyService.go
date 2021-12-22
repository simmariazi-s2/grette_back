package service

import (
	"log"
	"net/http"
	"work/grette_back/repositories"

	"github.com/gin-gonic/gin"
)

func GetCompanyList(c *gin.Context) {

	companyList, err := repositories.GetCompanyList()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    403,
			"message": "회사 목록 조회 실패",
		})

		log.Print(err.Error())

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "회사 목록 조회",
		"CompanyList": companyList,
	})
}
