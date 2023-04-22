package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn-gin/dao"
	"learn-gin/global"
	"learn-gin/models/api"
	"learn-gin/models/system"
)

type MenuService struct {
}

var menuDao = new(dao.MenuDao)

func (s MenuService) Get(c *gin.Context) {
	id := c.Param("id")
	menu := menuDao.GetById(id)
	api.Success(c, menu)
}

func (s MenuService) List(c *gin.Context) {
	var menus []*system.Menu
	global.DB.Find(&menus)
	result := recursiveNode(menus, "0")
	fmt.Println(result)
	api.Success(c, result)
}

func recursiveNode(list []*system.Menu, pId string) []*system.Menu {
	res := make([]*system.Menu, 0)

	for _, item := range list {
		if item.Pid == pId {
			item.Children = recursiveNode(list, item.Id)
			res = append(res, item)
		}
	}
	return res
}
