package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArticleController struct {
	BaseController
}

func (con ArticleController) Index(c *gin.Context) {
	//c.String(http.StatusOK, "文章列表")
	con.success(c)
}

func (con ArticleController) Add(c *gin.Context) {
	c.String(http.StatusOK, "添加文章")
}

func (con ArticleController) Edit(c *gin.Context) {
	c.String(http.StatusOK, "修改文章")
}
