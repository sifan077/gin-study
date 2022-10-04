package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

/**
*  中间件的使用
 */

func main() {
	// 创建一个不包含中间件的路由器
	router := gin.New()
	// 全局中间件
	// 使用 Logger 中间件
	router.Use(gin.Logger())

	// 使用 Recovery 中间件
	router.Use(gin.Recovery())

	// 给指定路由绑定中间件
	router.GET("/login", loginMid(), func(context *gin.Context) {
		context.String(http.StatusOK, "成功登陆")
	})

	// 给路由组绑定中间件
	v1 := router.Group("/v1")
	v1.Use(v1Mid())
	{
		v1.GET("/a", func(context *gin.Context) {
			context.String(http.StatusOK, "a")
		})
		v1.GET("/b", func(context *gin.Context) {
			context.String(http.StatusOK, "b")
		})
		v1.GET("/c", func(context *gin.Context) {
			context.String(http.StatusOK, "c")
		})
	}

	router.Run("127.0.0.1:8080")
}

func v1Mid() gin.HandlerFunc {
	return func(context *gin.Context) {
		log.Println("经过v1组中间件.............")
		v1 := context.DefaultQuery("v1", "")
		log.Println(v1)
		if v1 == "" {
			context.Abort()
			log.Println("不允许访问/v1时不带v1参数")
			context.String(http.StatusForbidden, "权限不足")
			return
		}
		context.Next()
	}
}

func loginMid() gin.HandlerFunc {
	// 构建自己的中间件的使用
	return func(context *gin.Context) {
		log.Println("经过login中间件.................")
		// 需要带着参数来访问login
		name := context.DefaultQuery("name", "")
		password := context.DefaultQuery("password", "")
		log.Println(name)
		log.Println(password)

		if name == "" || password == "" {
			context.Abort()
			log.Println("不允许访问/login时不带参数")
			context.String(http.StatusForbidden, "请先登录")
			return
		}
		context.Next()
	}
}
