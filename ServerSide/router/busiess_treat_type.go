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
		TreatTypeRouter.POST("getalls", business.GetAlls)
		TreatTypeRouter.POST("getbyid", business.GetById)
		TreatTypeRouter.POST("getpagelist", business.GetPageList)
	}
}