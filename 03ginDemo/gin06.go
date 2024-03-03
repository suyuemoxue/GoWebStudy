package main

//上传文件
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("03ginDemo/index.html")
	router.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})
	router.POST("/upload", func(context *gin.Context) {
		//从请求中读取文件
		file, err := context.FormFile("f1")
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			//保存到本地
			dst := path.Join("./", file.Filename)
			err := context.SaveUploadedFile(file, dst)
			if err != nil {
				return
			}
			context.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
	})
	router.Run(":8888")
}
