package apis

import (
	"github.com/gin-gonic/gin"
	"learn-gin/controllers/api"
	"net/http"
)

type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func UserRoutersInit(group *gin.RouterGroup) {
	userRouter := group.Group("/users")
	{
		userRouter.GET("", api.UserController{}.List)

		userRouter.GET("/:id", func(context *gin.Context) {
			id := context.Param("id")
			context.String(http.StatusOK, id)
		})

		userRouter.POST("/add", api.UserController{}.Add)
		userRouter.POST("/edit", api.UserController{}.Edit)
	}
}
