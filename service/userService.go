package service

import (
	"fmt"
	"log"
	"net/http"
	"work/grette_back/model"

	"github.com/gin-gonic/gin"
)

func SendEmail(c *gin.Context) {
	param := c.Param("param")

	c.JSON(http.StatusOK, gin.H{
		"mesasge": "메일발송",
		"param":   param,
	})
}

func CheckCode(c *gin.Context) {
	code := c.Query("code")

	c.JSON(http.StatusOK, gin.H{
		"message": "인증번호 체크",
		"id":      code,
	})
}

func CheckNickName(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "닉네임 체크",
	})
}

func CheckUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "로그인 & 유저 유효성 체크",
	})
}

func RegisterUser(c *gin.Context) {

	user := new(model.User)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    403,
			"message": "유저정보를 확인하세요",
		})

		log.Print(err.Error())

		return
	}

	fmt.Println("User :: ", user)

	c.JSON(http.StatusOK, gin.H{
		"Content-type": "application/json",
		"message":      "회원가입",
		"아이디":          user,
	})
}

func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "유저정보 조회",
	})
}

func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "회원 삭제",
	})
}

func UpdateUser(c *gin.Context) {

	user := new(model.User)
	fmt.Println("User1 :: ", user)
	c.BindJSON(&user)
	fmt.Println("User2 :: ", user.Email)
	fmt.Println("User2 :: ", user.Company)
	fmt.Println("User2 :: ", user.Password)
	c.JSON(http.StatusOK, gin.H{
		"message": "비밀번호 변경 및 닉네임 변경",
		"유저정보":    user,
	})
}
