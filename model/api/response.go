package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}

func Failed(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: 10000,
		Msg:  msg,
	})
}

func Forbidden(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: 10041,
		Msg:  msg,
	})
}
