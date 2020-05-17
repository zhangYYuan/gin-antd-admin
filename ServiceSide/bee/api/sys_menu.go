package api

import (
	"bee/model"
	"bee/service"
	"github.com/gin-gonic/gin"
)

func AddMenu(c *gin.Context)  {
	var menu model.SysMenu
	_ = c.ShouldBindJSON(&menu)
	err := service.AddMenu(menu)
	if err != nil {

	}
}