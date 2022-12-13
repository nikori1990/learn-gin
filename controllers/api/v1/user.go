package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn-gin/models"
	"net/http"
)

type UserController struct {
}

func (con UserController) Index(c *gin.Context) {
	//c.String(http.StatusOK, "用户列表")
	username := c.Query("username")
	age := c.Query("age")
	page := c.DefaultQuery("page", "1")

	c.JSON(http.StatusOK, gin.H{
		"username": username,
		"age":      age,
		"page":     page,
	})
}

func (con UserController) Add(c *gin.Context) {
	username, _ := c.Get("username")
	fmt.Println(username)

	user := &models.User{
		Username: "niko",
		Age:      22,
	}

	models.DB.Create(&user)
	fmt.Println(user)

	v, ok := username.(string)
	if ok {
		c.String(http.StatusOK, "添加用户 --- "+v)
	} else {
		c.String(http.StatusOK, "添加用户")
	}
}

func (con UserController) Edit(c *gin.Context) {
	c.String(http.StatusOK, "修改用户")
}

func (con UserController) List(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"result": users,
	})
}
