package service

import (
	"encoding/json"
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
	data, _ := json.Marshal(user)

	c.JSON(http.StatusOK, gin.H{
		"Content-type": "application/json",
		"message":      "회원가입",
		"아이디":          data,
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
	c.JSON(http.StatusOK, gin.H{
		"message": "비밀번호 변경 및 닉네임 변경",
	})
}
