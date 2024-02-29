package main

//获取query参数
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func query(context *gin.Context) {
	//name := context.Query("query") //获取请求中携带的参数
	//name := context.DefaultQuery("query", "墨雪") //获取请求中携带的参数，取不到就用默认值
	name, err := context.GetQuery("query")
	if !err {
		fmt.Println("获取失败")
		name = "墨雪"
	}
	context.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}

func main() {
	router := gin.Default()
	router.GET("/web", query)
	router.Run(":8888")
}
