package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn-gin/dao"
	"learn-gin/global"
	"learn-gin/models/system"
)

type MenuService struct {
}

var menuDao = new(dao.MenuDao)

func (s MenuService) Get(c *gin.Context) {
	id := c.Param("id")
	menu := menuDao.GetById(id)
	Success(c, menu)
}

func (s MenuService) List(c *gin.Context) {
	var menus []*system.SysMenu
	global.DB.Find(&menus)
	result := recursiveNode(menus, "0")
	fmt.Println(result)
	Success(c, result)
}

func recursiveNode(list []*system.SysMenu, pId string) []*system.SysMenu {
	res := make([]*system.SysMenu, 0)

	for _, item := range list {
		if item.Pid == pId {
			item.Children = recursiveNode(list, item.Id)
			res = append(res, item)
		}
	}
	return res
}
