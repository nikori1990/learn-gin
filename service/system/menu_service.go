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
	result := recursiveNode(menus, 0)
	fmt.Println(result)
	api.Success(c, result)
}

func recursiveNode(list []*system.Menu, pId uint) []*system.Menu {
	res := make([]*system.Menu, 0)

	for _, item := range list {
		if item.ParentId == pId {
			item.Children = recursiveNode(list, item.ID)
			res = append(res, item)
		}
	}
	return res
}
