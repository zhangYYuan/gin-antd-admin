package router

import (
	"github.com/gin-gonic/gin"
	"para.common/middleware"
)
func InitRouter() *gin.Engine {
	var Router = gin.Default();
	Router.Use(middleware.Cors())

	ApiGroup := Router.Group("") // 方便统一添加路由组前缀 多服务器上线使用

	InitUserRouter(ApiGroup)
	InitBaseRouter(ApiGroup)
	return Router
}