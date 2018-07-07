package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"test": "hoge"})
}

func main() {
	router := gin.Default()

	// API namespace
	v1 := router.Group("/api/v1")
	{
		v1.GET("/test", test)
	}

	router.Run(":8080")
}
