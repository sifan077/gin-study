package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

/**
*  使用日志
 */

func main() {
	router := gin.New()
	// 创建记录日志的文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要将日志同时写入文件和控制台，请使用以下代码
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	// LoggerWithFormatter 中间件会将日志写入 gin.DefaultWriter
	// 默认的是输出到控制台 gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 你的自定义格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "日志功能")
	})

	router.Run("127.0.0.1:8080")
}
