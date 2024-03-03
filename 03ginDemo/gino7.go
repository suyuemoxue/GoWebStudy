package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	//请求重定向
	router.GET("/baidu", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})
	//转发
	router.GET("/a", func(context *gin.Context) {
		context.Request.URL.Path = "/b"
		router.HandleContext(context)
	})

	router.GET("/b", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "b",
		})
	})

	router.Run(":8888")

}
