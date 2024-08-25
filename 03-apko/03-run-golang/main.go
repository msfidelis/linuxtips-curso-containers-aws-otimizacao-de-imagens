package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/linuxtips", func(c *gin.Context) {
		c.String(200, "VAAAAAAAAAAI")
	})

	r.Run(":8080")
}
