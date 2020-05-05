package main

import (
	"bee/router"
)

func main() {
	r := router.InitRouter()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}