package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn-gin/controllers/admin"
)

type LoginForm struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type LoginController struct {
	admin.BaseController
}

func (con LoginController) Login(c *gin.Context) {
	//form := LoginForm{}
	//err := c.BindJSON(&form)
	//if err != nil {
	//	panic("请求数据不合法")
	//}

	//fmt.Println("username:" + form.Username)
	//fmt.Println("password:" + form.Password)

	form := make(map[string]interface{})
	err := c.BindJSON(&form)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(form["username"])
	fmt.Println(form["password"])
}
