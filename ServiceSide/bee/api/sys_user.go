package api

import (
	"bee/global"
	"bee/global/response"
	"bee/middleware"
	"bee/model"
	"bee/model/request"
	"bee/service"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

// 用户注册
func Register(c *gin.Context) {
	var R request.RegisterResponse
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
		response.SuccessWithDescription(request.LoginResponse{User: userReturn}, "注册成功", c)
	}
}

// 用户登录
func Login(c *gin.Context) {
	var R request.RegisterAndLoginStruct
	_ = c.ShouldBindJSON(&R)

	u := &model.SysUser{UserName: R.Username, Password: R.Password}
	err, user := service.Login(u)
	if err != nil {
		response.FailWithDescription(response.ERROR, "", "用户名或密码错误", c)
	}  else {
		TokenNext(c, *user)
	}
}

func TokenNext(c *gin.Context, user model.SysUser)  {
	j := &middleware.JWT{
		[]byte(global.BeeConfig.JWT.SigningKey),
	}

	claims := request.CustomClaims{
		UUID: user.UUID,
		ID: user.ID,
		NickName: user.NickName,
		AuthorityId: user.AuthorityId,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),       // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 60*60*24*7), // 过期时间 一周
			Issuer:    "beeAdmin",                              //签名的发行者
		},
	}

	fmt.Println("------->>", claims)
	token, err := j.CreateToken(claims)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		if global.BeeConfig.System.UserMultiPoint {

		} else {
			response.SuccessWithData(request.LoginResponse{
				User: user,
				Token: token,
				ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
			}, c)
		}
	}

}