package routers

import (
	"log"
	"net/http"
	"time"
	"work/grette_back/service"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "없는페이지",
		})
	})

	router.Use(func(c *gin.Context) {
		start := time.Now()
		log.Print(`[START] `, c.FullPath(), ` START `, start.Format("2006-01-02 15:04:05"))
		c.Next()
		log.Print(`[END] `, c.FullPath(), ` END `, time.Since(start))
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Gin index",
		})
	})

	router.GET("/getCompanyList", service.GetCompanyList)

	userUrl := router.Group("/user")
	{
		// 이메일 발송
		userUrl.GET("/sendEmail", service.SendEmail)

		// 인증번호 체크
		userUrl.GET("/checkCode", service.CheckCode)

		// 닉네임 체크(중복 검사, 욕설)
		userUrl.GET("/checkNickName", service.CheckNickName)

		// 이메일 체크(중복 검사)
		userUrl.GET("/checkEmail", service.CheckEmail)

		// 회원가입
		userUrl.POST("/register", service.RegisterUser)

		// 유저정보 수정 (닉네임 변경,  비밀번호 변경)
		userUrl.POST("/update", service.UpdateUser)

		// 탈퇴
		userUrl.POST("/delete", service.DeleteUser)

		// 유저정보 조회
		userUrl.GET("/getInfo", service.GetUser)

		// 로그인
		userUrl.POST("/login", service.DoLogin)

	}

	return router
}
