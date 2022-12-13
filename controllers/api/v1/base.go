package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct {
}

func (con BaseController) success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (con BaseController) error(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
