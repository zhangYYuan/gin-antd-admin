package router

import (
	"gin-vue-admin/controller/business"
	"gin-vue-admin/middleware"
"github.com/gin-gonic/gin"
)

func InitTreatTypeRouter(Router *gin.RouterGroup) {
	TreatTypeRouter := Router.Group("treattype").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		TreatTypeRouter.POST("create", business.Create)
		TreatTypeRouter.PUT("update", business.Update)
		TreatTypeRouter.DELETE("delete", business.Delete)
		TreatTypeRouter.GET("getalls", business.GetAlls)
		TreatTypeRouter.GET("getbyid", business.GetById)
		TreatTypeRouter.GET("getpagelist", business.GetPageList)
	}
}