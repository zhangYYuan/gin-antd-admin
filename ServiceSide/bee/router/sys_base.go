package router

import (
	"bee/api"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes)  {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("register", api.Register)
		BaseRouter.POST("login", api.Login)
	}
	return BaseRouter
}
