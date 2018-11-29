package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{ //statuspost ,value
			"message": "hello world!",
		})
	})
	r.Run(":" + os.Getenv("PORT"))
}
