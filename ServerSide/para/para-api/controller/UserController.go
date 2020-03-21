package controller

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	models "para.common/models"
	"time"
)



func RegisterUser(c *gin.Context)  {
	user := models.User{}

	if len(user.Phone) != 11{
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg": "手机号必须为11位",
		})
		return
	}
	if err := c.BindJSON(&user); err != nil {
		
	}
	// 查库判断手机号是否存在
	//if ()
	c.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}

func LoginUser(c *gin.Context)  {
	var user models.User
	if err := c.BindJSON(&user); err != nil {

	}


}

func RandomString(n int) string{
	var letters = []byte("dfghjklertfghjkl")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}