package main

//获取URI路径参数
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/:name/:age", func(context *gin.Context) {
		name := context.Param("name")
		age := context.Param("age")
		context.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	router.Run(":8888")

}
