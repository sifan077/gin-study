package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

/**
* 自定义验证器
 */

type myPerson struct {
	//不能为空并且大于10
	Age int `form:"age" binding:"required,gt=10"`
	// 不能为空
	Name string `form:"name" binding:"required"`
	// 时间必须按照格式
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	router := gin.New()
	router.GET("checkin", func(context *gin.Context) {
		var person myPerson
		// 进行参数验证
		if err := context.ShouldBind(&person); err != nil {
			context.String(500, fmt.Sprint(err))
			return
		}
		context.String(200, fmt.Sprintf("%#v", person))
	})
	router.Run("127.0.0.1:8080")
}
