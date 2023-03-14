package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn-gin/middlewares"
	"learn-gin/models/api"
	"time"
)

type LoginForm struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type LoginService struct {
}

func (s LoginService) Login(c *gin.Context) {
	form := make(map[string]string)
	err := c.BindJSON(&form)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(form["username"])
	fmt.Println(form["password"])
	now := time.Now()
	expireTime := now.Add(time.Hour)
	token, err := middlewares.GenerateToken(form["username"], expireTime)
	if err != nil {
		fmt.Println("GenerateToken error")
	}
	fmt.Println("token:", token)

	api.Success(c, token)
}