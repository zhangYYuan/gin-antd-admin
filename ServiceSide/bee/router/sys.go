package router

import (
	"bee/api"
	"bee/middleware"
	"github.com/gin-gonic/gin"
)

func InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes)  {
	BaseRouter := Router.Group("menu").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		BaseRouter.POST("add", api.AddMenu)
	}
	return BaseRouter
}
