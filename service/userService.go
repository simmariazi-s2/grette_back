package service

import (
	"log"
	"net/http"
	"time"
	"work/grette_back/app"
	"work/grette_back/database/entities"
	"work/grette_back/message"
	"work/grette_back/model"
	"work/grette_back/repositories"
	"work/grette_back/util"

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

	if len(code) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    403,
			"message": "code 공백",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "인증번호 체크",
		"id":      code,
	})
}

func CheckNickName(c *gin.Context) {

	gin := app.Gin{C: c}

	nickName := c.Query("nickName")

	result, err := repositories.ExistsNickName(nickName)

	if err != nil {
		log.Print(err.Error())
		gin.Response(http.StatusBadRequest, message.INVALID_PARAMS, result)
		return
	}

	gin.Response(http.StatusOK, message.SUCCESS, result)
}

func CheckEmail(c *gin.Context) {

	gin := app.Gin{C: c}

	email := c.Query("email")
	result, err := repositories.ExistsUserEmail(email)

	if err != nil {
		gin.Response(http.StatusBadRequest, message.INVALID_PARAMS, result)
		return
	}

	gin.Response(http.StatusOK, message.SUCCESS, result)
}

func CheckUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "로그인 & 유저 유효성 체크",
	})
}

func RegisterUser(c *gin.Context) {

	gin := app.Gin{C: c}

	user := new(model.User)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    403,
			"message": "유저정보를 확인하세요",
		})

		log.Print(err.Error())

		return
	}

	newUser := new(entities.User)

	newUser.UserNickname = user.NickName
	newUser.CompanyNo = user.Company
	newUser.CreateDtm = (*time.Time)(time.Now().UTC().Location())

	result, err := repositories.CreateUser(*newUser)

	if err != nil {
		gin.Response(http.StatusBadRequest, message.CREATE_FAIL, user)
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"code":    403,
		// 	"message": "회원가입 에러",
		// 	"result":  result,
		// })

		// log.Print(err.Error())

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Content-type": "application/json",
		"message":      "회원가입",
		"이메일":          user.Email,
		"result":       result,
	})
}

func GetUser(c *gin.Context) {

	id := c.Query("id")

	user, err := repositories.GetUser(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    403,
			"message": "유저정보 조회 오류",
		})

		log.Print(err.Error())

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "유저정보 조회",
		"유저 정보":   user,
	})
}

func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "회원 삭제",
	})
}

func UpdateUser(c *gin.Context) {

	gin := app.Gin{C: c}

	user := new(model.User)

	c.BindJSON(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "비밀번호 변경 및 닉네임 변경",
		"유저정보":    user,
	})

	// if err != nil {
	// 	log.Print(err.Error())
	// 	gin.Response(http.StatusBadRequest, message.INVALID_PARAMS, result)
	// 	return
	// }

	// gin.Response(http.StatusOK, message.SUCCESS, result)

	userEntity := new(entities.User)
	userEntity.UserNickname = user.NickName

	//복호화 추가 필요
	// 다시 인코딩
	encodePassword := util.EncodeBase64(user.Password)
	userEntity.UserPassword = encodePassword
	userEntity.UpdateDtm = time.Now()
	userEntity.CompanyNo = user.Company

	result, err := repositories.UpdateUser(*userEntity)

	if err != nil {
		log.Print(err.Error())
		gin.Response(http.StatusBadRequest, message.UPDATE_FAIL, result)
		return
	}

	gin.Response(http.StatusOK, message.SUCCESS, result)
}

func DoLogin(c *gin.Context) {
	gin := app.Gin{C: c}

	user := new(model.User)

	c.BindJSON(&user)

	if len(user.Email) == 0 || len(user.Password) == 0 {

		gin.Response(http.StatusBadRequest, message.ERROR, user)

		return
	}

	result, err := repositories.ExistsPassword(user.Email, user.Password)

	if err != nil {

		gin.Response(http.StatusBadRequest, message.LOGIN_FAIL, result)
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"code":    403,
		// 	"message": "로그인 오류",
		// })

		// log.Print(err.Error())

		return
	}

	gin.Response(http.StatusOK, message.SUCCESS, result)
}
