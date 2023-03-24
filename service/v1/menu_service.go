package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn-gin/global"
	"learn-gin/models"
	"learn-gin/models/api"
)

type MenuService struct {
}

func (s MenuService) Get(c *gin.Context) {
	id := c.Param("id")
	var menu *models.Menu
	global.DB.Where("id=?", id).First(&menu)
	api.Success(c, menu)
}

func (s MenuService) List(c *gin.Context) {
	var menus []*models.Menu
	global.DB.Find(&menus)
	result := recursiveNode(menus, "0")
	fmt.Println(result)
	api.Success(c, result)
}

func recursiveNode(list []*models.Menu, pId string) []*models.Menu {
	res := make([]*models.Menu, 0)

	for _, item := range list {
		if item.Pid == pId {
			item.Children = recursiveNode(list, item.Id)
			res = append(res, item)
		}
	}
	return res
}
