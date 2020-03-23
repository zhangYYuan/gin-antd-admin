package sys

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

func GetCurrentUser(c *gin.Context)  {
	fmt.Println("ddd")
}

func RegisterUser(c *gin.Context)  {
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