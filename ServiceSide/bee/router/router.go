package router
import (
"github.com/gin-gonic/gin"
)
func InitRouter() *gin.Engine {
	var Router = gin.Default();
	ApiGroup := Router.Group("") // 方便统一添加路由组前缀 多服务器上线使用
	InitBaseRouter(ApiGroup)
	return Router
}