package api

import (
	"bee/model"
	"bee/model/request"
	"bee/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var R request.ResisterStruct
	_ = c.ShouldBindJSON(&R)
	user := &model.SysUser{
		UserName: R.Username,
		NickName: R.NickName,
		Password: R.Password,
		HeaderImg: R.HeaderImg,
	}
	err, userReturn := service.Resigster(*user)
	fmt.Println(err, userReturn)
}