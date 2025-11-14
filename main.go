package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hi Rezvan!"})
	})
	r.Run(":8080")
	fmt.Print("Hello Go !")
}
