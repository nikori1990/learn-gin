package v1

import (
	"github.com/gin-gonic/gin"
	"learn-gin/controllers/api/v1"
	"net/http"
)

type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func UserRoutersInit(group *gin.RouterGroup) {
	userRouter := group.Group("/users")
	{
		userRouter.GET("", v1.UserController{}.List)

		userRouter.GET("/:id", func(context *gin.Context) {
			id := context.Param("id")
			context.String(http.StatusOK, id)
		})

		userRouter.POST("/add", v1.UserController{}.Add)
		userRouter.POST("/edit", v1.UserController{}.Edit)
	}
}
