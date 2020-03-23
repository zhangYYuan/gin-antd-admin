package main

import (
	"github.com/gin-gonic/gin"
	_ "para/para.common/db"
)


func main() {
	r := gin.Default()

	r.Run(":3000")
}