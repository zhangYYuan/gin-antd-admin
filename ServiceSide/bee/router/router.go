package router
import (
	"bee/middleware"
	"github.com/gin-gonic/gin"
)
func InitRouter() *gin.Engine {
	var Router = gin.Default();
	Router.Use(middleware.Cors())


	ApiGroup := Router.Group("/hb") // 方便统一添加路由组前缀 多服务器上线使用
	InitBaseRouter(ApiGroup)


	return Router
}