package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("user/search", func(c *gin.Context) {
		name := c.Query("name")
		address := c.Query("address")
		c.JSON(http.StatusOK, gin.H{
			"name":    name,
			"address": address,
		})
	})

	r.GET("user/search/:name/:address", func(c *gin.Context) {
		name := c.Param("name")
		address := c.Param("address")
		c.JSON(http.StatusOK, gin.H{
			"name":    name,
			"address": address,
		})
	})
	r.Run(":8000")
}
