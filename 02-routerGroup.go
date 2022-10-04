package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
* 路由分组
 */

func main() {
	//router := gin.Default()
	// 不使用默认中间件创建gin路由器
	router := gin.New()
	// 分组一
	v1 := router.Group("/v1")
	{
		v1.GET("/login", func(context *gin.Context) {
			context.String(http.StatusOK, "v1/login")
		})
		v1.GET("/submit", func(context *gin.Context) {
			context.String(http.StatusOK, "v1/submit")
		})
		v1.GET("/read", func(context *gin.Context) {
			context.String(http.StatusOK, "v1/read")
		})
	}
	// 分组二
	v2 := router.Group("/v2")
	{
		v2.GET("/login", func(context *gin.Context) {
			context.String(http.StatusOK, "v2/login")
		})
		v2.GET("/submit", func(context *gin.Context) {
			context.String(http.StatusOK, "v2/submit")
		})
		v2.GET("/read", func(context *gin.Context) {
			context.String(http.StatusOK, "v2/read")
		})
	}

	router.Run("127.0.0.1:8080")
}
