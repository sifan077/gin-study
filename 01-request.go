package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"log"
	"net/http"
	"time"
)

/**
 * 各种请求的使用
 */

// StatCost 中间件需要返回 gin.HandlerFunc
func StatCost() gin.HandlerFunc {
	// 构建自己的中间件的使用
	return func(c *gin.Context) {
		log.Println("经过中间件.................")

		// 不允许访问 /c
		// 获取路径并且比较
		flag := c.FullPath() == "/c"
		if flag {
			c.Abort()
			log.Println("不允许访问/c")
			c.String(http.StatusForbidden, "不允许访问")
		}

		start := time.Now()
		//送往下一级的数据
		//c.Set("name","hello")
		//送往下一级
		c.Next()
		//阻止往下传递
		//c.Abort()
		cost := time.Since(start)
		log.Println(cost)
	}
}

func main() {
	// 禁用控制台颜色
	gin.DisableConsoleColor()

	// 使用默认中间件创建一个gin路由器
	// logger and recovery (crash-free) 中间件
	router := gin.Default()

	// 不使用默认中间件创建gin路由器
	//router:=gin.New()

	// 注册自己的中间件
	router.Use(StatCost())

	// 普通get请求
	router.GET("/someGet", getting)
	// 带参数的get请求
	router.GET("/getPara", getPara)
	// 提交表单的post请求
	router.POST("/somePost", posting)
	// 带body参数的post请求
	router.POST("/postBody", postBody)
	// url传递参数
	router.GET("/user/search/:username/:age", urlPara)

	//重定向
	router.GET("/index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	//请求转发
	router.GET("/a", func(c *gin.Context) {
		c.Request.URL.Path = "/b"
		router.HandleContext(c)
	})
	router.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	router.GET("/c", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// 以下各种请求都是跟post差不多的用法
	//router.PUT("/somePut", putting)
	//router.DELETE("/someDelete", deleting)
	//router.PATCH("/somePatch", patching)
	//router.HEAD("/someHead", head)
	//router.OPTIONS("/someOptions", options)

	// 默认启动的是 8080端口，也可以自己定义启动端口
	router.Run("127.0.0.1:8080")
	// router.Run(":3000") for a hard coded port
}

func urlPara(context *gin.Context) {
	// url传参
	username := context.Param("username")
	age := context.Param("age")
	log.Println(username)
	log.Println(age)
	context.String(http.StatusOK, "Success")
}

func postBody(context *gin.Context) {
	// post 获取提交的body
	// 不处理错误了
	body, _ := context.GetRawData()
	// 定义map或结构体
	var m map[string]interface{}
	// 反序列化
	json.Unmarshal(body, &m)
	log.Println(m)
	context.JSON(http.StatusOK, m)
}

func posting(context *gin.Context) {
	// post 获取提交的表单
	name := context.PostForm("name")
	fmt.Println(name)
	context.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}

func getPara(context *gin.Context) {
	/**
	  第一个参数为key
	  第二个参数代表如果没有这个参数 使用这个作为默认值
	*/
	name := context.DefaultQuery("name", "哈哈")

	age := context.Query("age")

	context.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}

func getting(context *gin.Context) {
	// get请求返回json
	context.JSON(http.StatusOK, gin.H{
		"name": "shentu",
		"age":  123,
	})
}
