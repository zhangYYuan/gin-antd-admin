package sys

import (
	"crypto/dsa"
	"github.com/gin-gonic/gin"
	"net/http"
	"para.common/models"
)

func Login(c * gin.Context)  {

	var user models.SysUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := Genera
	//fmt.Println("Login--->", user)

}