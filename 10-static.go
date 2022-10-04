package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
* 设置静态文件
 */

func main() {
	router := gin.Default()
	// 将assets 作为访问的路径，其实际文件在 ./assets下
	router.Static("/assets", "./assets")
	// 将一个文件夹映射出去
	router.StaticFS("/css", http.Dir("./assets/css"))
	router.StaticFS("/js", http.Dir("./assets/js"))
	// 映射一个静态文件文件
	router.StaticFile("/favicon.ico", "./assets/favicon.ico")

	router.Run("127.0.0.1:8080")
}
