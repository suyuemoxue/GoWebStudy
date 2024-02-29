package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("03template/hello.gohtml")
	r.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "hello.gohtml", gin.H{
			"Name":   "tom",
			"Gender": "男",
			"Age":    20,
		})
	})
	r.Run(":8888")
}
