package system

import (
	"github.com/gin-gonic/gin"
	"learn-gin/model/api"
	"learn-gin/model/system"
	"learn-gin/model/system/query"
	"strconv"
)

type UserService struct {
}

func (UserService) Create(c *gin.Context) {
	var user system.User
	if err := c.ShouldBindJSON(&user); err != nil {
		panic(err)
	}
	createId := userRepository.Create(&user)
	api.Success(c, createId)
}

func (UserService) Update(c *gin.Context) {
	var user system.User
	if err := c.ShouldBindJSON(&user); err != nil {
		panic(err)
	}
	updateId := userRepository.Update(&user)
	api.Success(c, updateId)
}

func (UserService) Delete(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		panic(err)
	}
	deleteId := userRepository.Delete(uint(id))
	api.Success(c, deleteId)
}

func (UserService) GetById(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		panic(err)
	}
	user := userRepository.GetById(uint(id))
	api.Success(c, user)
}

func (UserService) List(c *gin.Context) {
	list := userRepository.List()
	api.Success(c, list)
}

func (UserService) Page(c *gin.Context) {
	var pageQuery query.UserQuery

	if err := c.ShouldBindQuery(&pageQuery); err != nil {
		panic(err)
	}

	pageQuery.SetFromLimit()
	page := userRepository.Page(&pageQuery)

	api.Success(c, page)
}
