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

func (router *UserRouter) Init(group *gin.RouterGroup) {
	userRouter := group.Group("/users")
	{
		userRouter.GET("", v1.UserService{}.List)

		userRouter.GET("/:id", func(context *gin.Context) {
			id := context.Param("id")
			context.String(http.StatusOK, id)
		})

		userRouter.POST("/add", v1.UserService{}.Add)
		userRouter.POST("/edit", v1.UserService{}.Edit)
	}
}
