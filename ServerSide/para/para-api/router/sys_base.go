package router

import (
	"github.com/gin-gonic/gin"
	"para/para-api/controller/sys"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("login", sys.Login)

	}
	return BaseRouter
}
