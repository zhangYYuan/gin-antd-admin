package middleware

import (
	"bee/global/response"
	"bee/model/request"
	"bee/service"
	"github.com/gin-gonic/gin"
)

func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		waitUser := claims.(*request.CustomClaims)

		// 获取请求的URI
		obj := c.Request.URL.RequestURI()

		// 获取请求的方法
		act := c.Request.Method

		// 获取用户的角色
		sub := waitUser.AuthorityId

		e := service.Casbin()

		if e.Enforce(sub, obj, act) {
			c.Next()
		} else {
			response.Result(response.ERROR, gin.H{}, "权限不足", c)
			c.Abort()
			return
		}
	}
}
