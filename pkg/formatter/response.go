package app

import (
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, msg string, data interface{}) {
	g.C.Header("Content-Type", "application/json")
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  msg,
		// Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}
