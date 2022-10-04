package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
*  响应不同的数据
 */

func main() {
	router := gin.Default()
	router.GET("/xml", func(context *gin.Context) {
		context.XML(http.StatusOK, gin.H{"message": "hello", "status": http.StatusOK})
	})
	router.GET("/json", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "hello", "status": http.StatusOK})
	})
	router.GET("/yaml", func(context *gin.Context) {
		context.YAML(http.StatusOK, gin.H{"message": "hello", "status": http.StatusOK})
	})
	//router.GET("/protoBuf", func(context *gin.Context) {
	//	reps := []int64{int64(1), int64(2)}
	//	label := "test"
	//	// The specific definition of protobuf is written in the testdata/protoexample file.
	//	data := &protoexample.Test{
	//		Label: &label,
	//		Reps:  reps,
	//	}
	//	// Note that data becomes binary data in the response
	//	// Will output protoexample.Test protobuf serialized data
	//	context.ProtoBuf(http.StatusOK, data)
	//})

	router.Run("127.0.0.1:8080")
}
