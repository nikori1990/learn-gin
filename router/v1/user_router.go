package v1

import (
	"github.com/gin-gonic/gin"
	"learn-gin/service/v1"
	"net/http"
)

type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type UserRouter struct {
}

var userService = new(v1.UserService)

func (router *UserRouter) Init(group *gin.RouterGroup) {
	userRouter := group.Group("/users")
	{
		userRouter.GET("", userService.List)

		userRouter.GET("/:id", func(context *gin.Context) {
			id := context.Param("id")
			context.String(http.StatusOK, id)
		})

		userRouter.POST("/add", userService.Add)
		userRouter.POST("/edit", userService.Edit)
	}
}
