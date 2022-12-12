package api

import (
	"github.com/gin-gonic/gin"
	"learn-gin/models"
	"learn-gin/models/api"
)

type MenuController struct {
}

func (con MenuController) Get(c *gin.Context) {
	id := c.Param("id")
	var menu *models.Menu
	models.DB.Where("id=?", id).First(&menu)
	api.Success(c, menu)
}

func (con MenuController) List(c *gin.Context) {
	var menus []*models.Menu
	models.DB.Find(&menus)
	result := recursiveNode(menus, "0")
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
