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

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	appGin := app.Gin{C: c}
	id := c.Query("id")

	user, err := repositories.GetUser(id)

	if err != nil {
		appGin.Response(http.StatusBadRequest, message.ERROR, user)
		log.Print(err.Error())
		return

	}

	appGin.Response(http.StatusOK, message.SUCCESS, user)
	log.Print("유저정보 조회 성공")

}

// 게시글 목록 조회
func SelectBoardList(c *gin.Context) {
	gin := app.Gin{C: c}

	boardList, err := repositories.GetBoardList()

	if err != nil {
		gin.Response(http.StatusBadRequest, message.ERROR, boardList)
		log.Print(err.Error())
		return
	}

	gin.Response(http.StatusOK, message.SUCCESS, boardList)
	log.Print("게시글 목록 조회 성공")
}

// 게시글 조회
func SelectBoard(c *gin.Context) {

}

// 게시글 쓰기
func RegisterBoard(c *gin.Context) {

}

// 게시글 수정
func UpdateBoard(c *gin.Context) {

}

// 게시글 삭제
func DeleteBoard(c *gin.Context) {

}

// 댓글 목록 조회
func SelectReplyList(c *gin.Context) {

}

// 댓글 조회
func SelectReply(c *gin.Context) {

}

// 댓글 쓰기
func RegisterReply(c *gin.Context) {

}

// 댓글 수정
func UpdateReply(c *gin.Context) {

}

// 댓글 쓰기
func DeleteReply(c *gin.Context) {

}

// 좋아요 (게시글 or 댓글)
func SetRecommand(c *gin.Context) {

}

// 추가 예정
func SendEmail(c *gin.Context) {
	param := c.Param("param")

	c.JSON(http.StatusOK, gin.H{
		"mesasge": "메일발송",
		"param":   param,
	})
}

// 코드정보 유무 체크
func CheckCode(c *gin.Context) {
	gin := app.Gin{C: c}
	code := c.Query("code")

	if len(code) == 0 {

		gin.Response(http.StatusBadRequest, message.INVALID_PARAMS, code)
		return
	}

	gin.Response(http.StatusOK, message.SUCCESS, code)
}

// 닉네임 존재 유무 체크
func ExistsNickName(c *gin.Context) {
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

// 회원 이메일 존재 유무 체크
// 성공시: row 반환, 실패: 0
func ExistsEmail(c *gin.Context) {

	gin := app.Gin{C: c}

	email := c.Query("email")
	result, err := repositories.ExistsUserEmail(email)

	if err != nil {
		gin.Response(http.StatusBadRequest, message.INVALID_PARAMS, result)
		log.Print(err.Error())
		return
	}

	gin.Response(http.StatusOK, message.SUCCESS, result)
}

// 회원가입
// 성공: row integer, 실패: 0
func RegisterUser(c *gin.Context) {

	appGin := app.Gin{C: c}

	user := new(model.User)

	if err := c.ShouldBindJSON(&user); err != nil {

		appGin.Response(http.StatusBadRequest, message.CREATE_FAIL, user)

		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"code":    403,
		// 	"message": "유저정보를 확인하세요",
		// })

		log.Print(err.Error())

		return
	}
	newUser := new(entities.User)
	newUser.UserNickname = &user.NickName
	newUser.CompanyNo = &user.Company

	result, err := repositories.CreateUser(*newUser)

	if err != nil {
		appGin.Response(http.StatusBadRequest, message.CREATE_FAIL, user)
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"code":    403,
		// 	"message": "회원가입 에러",
		// 	"result":  result,
		// })

		// log.Print(err.Error())

		return
	}

	appGin.Response(http.StatusOK, message.SUCCESS, result)
	log.Print("회원가입 성공")

	// c.JSON(http.StatusOK, gin.H{
	// 	"Content-type": "application/json",
	// 	"message":      "회원가입",
	// 	"이메일":          user.Email,
	// 	"result":       result,
	// })
}

func GetUser(c *gin.Context) {
	appGin := app.Gin{C: c}
	id := c.Query("id")

	user, err := repositories.GetUser(id)

	if err != nil {
		appGin.Response(http.StatusBadRequest, message.ERROR, user)
		log.Print(err.Error())
		return
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"code":    403,
		// 	"message": "유저정보 조회 오류",
		// })

		// log.Print(err.Error())

		// return
	}

	appGin.Response(http.StatusOK, message.SUCCESS, user)
	log.Print("유저정보 조회 성공")

	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "유저정보 조회",
	// 	"유저 정보":   user,
	// })
}

// 회원 정보 삭제
func DeleteUser(c *gin.Context) {
	appGin := app.Gin{C: c}
	user := new(model.User)

	if err := c.ShouldBindJSON(&user); err != nil {
		appGin.Response(http.StatusBadRequest, message.ERROR, user)
		return
	}

	dbUser := new(entities.User)

	dbUser.UserNo = user.UserNo
	dbUser.IsUsed = 0

	repositories.UpdateUser(*dbUser)

	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "회원 삭제",
	// })
}

// 회원 정보 수정
func UpdateUser(c *gin.Context) {

	gin := app.Gin{C: c}

	user := new(model.User)

	if err := c.ShouldBindJSON(&user); err != nil {
		gin.Response(http.StatusBadRequest, message.INVALID_PARAMS, user)
		log.Print(err.Error())
		return
	}

	// c.BindJSON(&user)

	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "비밀번호 변경 및 닉네임 변경",
	// 	"유저정보":    user,
	// })

	// if err != nil {
	// 	log.Print(err.Error())
	// 	gin.Response(http.StatusBadRequest, message.INVALID_PARAMS, result)
	// 	return
	// }

	// gin.Response(http.StatusOK, message.SUCCESS, result)

	userEntity := new(entities.User)
	userEntity.UserNickname = &user.NickName
	userEntity.UserPassword = &user.Password
	userEntity.UpdateDtm = time.Now()
	userEntity.CompanyNo = &user.Company

	result, err := repositories.UpdateUser(*userEntity)

	if err != nil {
		log.Print(err.Error())
		gin.Response(http.StatusBadRequest, message.UPDATE_FAIL, result)
		return
	}

	gin.Response(http.StatusOK, message.SUCCESS, result)
	log.Print("비밀번호 및 닉네임 변경 성공")
}

// 회원정보 로그인
func DoLogin(c *gin.Context) {
	gin := app.Gin{C: c}

	user := new(model.User)

	//c.BindJSON(&user)

	// user정보 체크
	if err := c.ShouldBindJSON(&user); err != nil {
		gin.Response(http.StatusBadRequest, message.INVALID_PARAMS, user)
		return
	}

	// user의 이메일/비밀번호 체크
	if len(user.Email) == 0 || len(user.Password) == 0 {
		gin.Response(http.StatusBadRequest, message.ERROR, user)
		return
	}

	// 비밀번호 체크
	result, err := repositories.ExistsPassword(user.Email, user.Password)

	if err != nil {
		gin.Response(http.StatusBadRequest, message.LOGIN_FAIL, result)
		return
	}

	gin.Response(http.StatusOK, message.SUCCESS, result)
}
