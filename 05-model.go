package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
*  参数与模型结构体绑定
*  关于这个把结构体创建的变量和参数绑定起来，可以把接收的参数赋给变量，有种sprig mvc直接使用对象接收参数的意思
 */

// myUser 声明一个结构体实体类
// 需要把结构体内的变量名首字母大写，这样才能访问到
type myUser struct {
	// 需要name参数  是字符串
	Name string `form:"name" json:"name" xml:"name"  binding:"required"`
	// 需要 password参数，是字符串
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	//  绑定json参数
	router.POST("/login/json", func(context *gin.Context) {
		var user myUser
		if err := context.ShouldBindJSON(&user); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if user.Name != "shentu" || user.Password != "1234" {
			context.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})
	//  绑定xml参数
	router.POST("/login/xml", func(c *gin.Context) {
		var user myUser
		if err := c.ShouldBindXML(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if user.Name != "shentu" || user.Password != "1235" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})
	//  绑定form参数
	router.POST("/login/form", func(c *gin.Context) {
		var user myUser
		// This will infer what binder to use depending on the content-type header.
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if user.Name != "shentu" || user.Password != "1234" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})
	router.Run("127.0.0.1:8080")
}
