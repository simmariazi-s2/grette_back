package app

import (
	"work/grette_back/message"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code:    errCode,
		Message: message.GetMessage(errCode),
		Data:    data,
	})
}
