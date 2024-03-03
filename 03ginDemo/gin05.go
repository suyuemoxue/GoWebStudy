package main

//参数绑定
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	UserName string `form:"username"`
	Password string `form:"password"`
}

func main() {
	router := gin.Default()
	router.GET("/user", func(context *gin.Context) {
		var user UserInfo
		err := context.ShouldBind(&user)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"message": user,
			})
		}
	})
	router.Run(":8888")

}
