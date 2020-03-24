package sys

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"para.common/conf"
	"para.common/middleware"
	"para.common/models"
)

func Login(c * gin.Context)  {

	var user models.SysUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Login--->", conf.GlobalConfig.JWT.SigningKey)

	j :=  &middleware.JWT{SigningKey:[]byte(conf.GlobalConfig.JWT.SigningKey)} // 唯一签名

	token, err := j.GenerateToken(user)
	if (err  != nil) {
		fmt.Println("Login--->", err)
	}
	fmt.Println(token)
}



