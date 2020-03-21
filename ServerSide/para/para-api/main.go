package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"para/para-api/controller"
)


func main() {
	r := gin.Default()

	r.Use(cors.Default())
	r.POST("/register", controller.RegisterUser)
	r.POST("/login", controller.LoginUser)
	r.Run(":3000")
}