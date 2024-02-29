package main

// 获取form表单参数
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func login(context *gin.Context) {
	context.HTML(http.StatusOK, "login.html", nil)
}

func welcome(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")
	context.HTML(http.StatusOK, "welcome.html", gin.H{
		"username": username,
		"password": password,
	})

}

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("03ginDemo/login.html", "03ginDemo/welcome.html")
	router.GET("/login", login)
	router.POST("/login", welcome)
	router.Run(":8888")
}
