package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"para/para-api/controller/sys"
)

func InitUserRouter(r *gin.RouterGroup)  {
	fmt.Println("ddd")
	UserRouter := r.Group("user").Use()
	{
		UserRouter.GET("currentUser", sys.GetCurrentUser)
	}
}