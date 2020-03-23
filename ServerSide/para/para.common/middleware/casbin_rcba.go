package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func Auth() gin.HandlerFunc  {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == '' || strings.HasPrefix(token, "Bearer") {

		}

		tokenString = token
	}
}

