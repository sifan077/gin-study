package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

/*
*  cookie的使用
 */

func main() {

	router := gin.Default()

	router.GET("/cookie", func(c *gin.Context) {

		// 获取cookie
		cookie, err := c.Cookie("gin_cookie")
		// 如果不存在
		if err != nil {
			// 设置一个cookie
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "default cookie",
				3600, "/", "localhost:8080", false, true)
			c.String(http.StatusForbidden, "不存在cookie")
		}
		c.String(http.StatusOK, cookie)
		log.Printf("Cookie value: %s \n", cookie)
	})

	router.Run("localhost:8080")
}
