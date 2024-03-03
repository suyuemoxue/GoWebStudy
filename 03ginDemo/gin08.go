package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func indexHandler(context *gin.Context) {
	fmt.Println("indexHandler start")
	context.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
	fmt.Println("indexHandler end")
}

// 中间件函数，用于统计耗时
func costTime(context *gin.Context) {
	fmt.Println("costTime start")
	start := time.Now() //开始计时
	context.Next()      //调用后面的函数
	//context.Abort()     //阻止执行后面的函数
	fmt.Println("costTime end")
	cost := time.Since(start) //结束计时
	fmt.Println("耗时:", cost)
}

func main() {
	router := gin.Default()
	router.Use(costTime) //全局注册中间件函数
	router.GET("/handle", indexHandler)
	router.Run()
}
