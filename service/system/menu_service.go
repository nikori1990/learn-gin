package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn-gin/global"
	"learn-gin/model/api"
	"learn-gin/model/system"
)

type MenuService struct {
}

func (MenuService) Get(c *gin.Context) {
	id := c.Param("id")
	menu := menuRepository.GetById(id)
	api.Success(c, menu)
}

func (MenuService) List(c *gin.Context) {
	var menus []*system.Menu
	global.DB.Find(&menus)
	result := recursiveMenu(menus, 0)
	fmt.Println(result)
	api.Success(c, result)
}

func recursiveMenu(list []*system.Menu, pId uint) []*system.Menu {
	res := make([]*system.Menu, 0)

	for _, item := range list {
		if item.ParentId == pId {
			item.Children = recursiveMenu(list, item.ID)
			res = append(res, item)
		}
	}
	return res
}
