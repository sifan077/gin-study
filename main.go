package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
* gin  第一个HelloWorld
 */
func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello shentu ! This is your first Gin project: gin-study",
		})
	})
	fmt.Println("启动在===> http://127.0.0.1:8080")
	router.Run("127.0.0.1:8080")
}
