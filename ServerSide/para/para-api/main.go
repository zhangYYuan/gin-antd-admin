package main

import (
	_ "para.common/db"
	"para/para-api/router"
)


func main() {
	r := router.InitRouter()
	r.Run(":3000")
}