package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type NewsController struct {
	BaseController
}

func (con NewsController) Index(c *gin.Context) {
	//c.String(http.StatusOK, "文章列表")
	con.success(c)
}

func (con NewsController) Add(c *gin.Context) {
	c.String(http.StatusOK, "添加news")
}

func (con NewsController) Edit(c *gin.Context) {
	c.String(http.StatusOK, "修改news")
}
