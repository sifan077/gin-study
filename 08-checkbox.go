package main

import "github.com/gin-gonic/gin"

/**
* 获取checkbox的参数,
* 获取多选框的多个参数
 */

type mySelect struct {
	Colors []string `form:"colors[]"`
}

func formHandler(c *gin.Context) {
	var fakeForm mySelect
	c.ShouldBind(&fakeForm)
	c.JSON(200, gin.H{"color": fakeForm.Colors})
}

func main() {
	router := gin.New()
	router.GET("/checkbox", formHandler)
	router.Run("127.0.0.1:8080")
}
