package app

import (
	"work/grette_back/message"

	"github.com/gin-gonic/gin"
)

// Gin.Context 타입을 가지는 구조체
type Gin struct {
	C *gin.Context
}

// Response시 필요한 항목 구조체 정의
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:data`
}

// Gin구조체의 gin.context로 설정된 gin Response 항수 리시버
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code:    errCode,
		Message: message.GetMessage(errCode),
		Data:    data,
	})
}
