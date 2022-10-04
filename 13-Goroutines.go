package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

/**
* 在中间件或处理程序中启动新的Goroutines时，
* 你不应该使用其中的原始上下文，你必须使用只读副本
* gin框架的每一个request请求都会开一个goroutine来处理，所以天然支持高并发
 */

func main() {
	r := gin.Default()

	r.GET("/long_async", func(c *gin.Context) {
		// 创建要在goroutine中使用的副本
		cCp := c.Copy()
		go func() {
			// simulate a long task with time.Sleep(). 2 seconds
			time.Sleep(2 * time.Second)
			// 这里使用你创建的副本
			// 异步执行，前面的睡眠结束才会执行到这
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
	})

	r.GET("/long_sync", func(c *gin.Context) {
		// simulate a long task with time.Sleep(). 2 seconds
		time.Sleep(2 * time.Second)
		// 这里没有使用goroutine，所以不用使用副本
		// 同步执行，睡眠和此句一同执行
		log.Println("Done! in path " + c.Request.URL.Path)
	})

	r.Run("localhost:8080")
}
