package api

import (
	"bee/common/response"
	"bee/model"
	"bee/model/request"
	"bee/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
  注册用户
 */
func Register(c *gin.Context) {
	var R request.RegisterStruct
	_ = c.ShouldBindJSON(&R)
	user := &model.SysUser{
		UserName: R.Username,
		NickName: R.NickName,
		Password: R.Password,
		HeaderImg: R.HeaderImg,
	}
	err, userReturn := service.Register(*user)
	if err != nil {
		response.FailWithDescription(response.ERROR, "", fmt.Sprintf("%v", err), c)
	} else {
		response.SuccessWithDescription(request.SysUserResponse{User: userReturn}, "注册成功", c)
	}
	fmt.Println(err, userReturn)
}