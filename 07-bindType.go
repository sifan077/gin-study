package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

/**
* ShouldBindQuery 函数只绑定Get参数，不绑定post数据
* ShouldBind 函数绑定Post参数
 */

type myHome struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	router := gin.New()
	router.Any("/test1", startPage)
	router.Any("/test2", endPage)

	router.POST("/postBind", endPage)
	router.Run("127.0.0.1:8080")
}

func startPage(c *gin.Context) {
	var person myHome
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "Success")
}

func endPage(c *gin.Context) {
	var person myHome
	// If `GET`, only `Form` binding engine (`query`) used.
	// 如果是Get，那么接收不到请求中的Post的数据？？
	// 如果是Post, 首先判断 `content-type` 的类型 `JSON` or `XML`, 然后使用对应的绑定器获取数据.
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
	}

	c.String(200, "Success")
}
