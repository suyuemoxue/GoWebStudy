package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(context *gin.Context) {
	user := User{
		Name:   "墨雪",
		Gender: "男",
		Age:    18,
	}

	context.JSON(http.StatusOK, user)
}

func main() {
	r := gin.Default() //返回默认的路由引擎
	//用户使用GET请求访问/hello时，执行sayHello函数
	r.GET("/hello", sayHello)
	r.Run(":8888")
}
