package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
)

func main() {
	r := gin.Default()
	// 处理multipart forms提交文件时默认内存限制是32MiB
	// 下列方式修改
	// r.MaxMultipartMemory = 8 << 20
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/upload", func(c *gin.Context) {
		// 从请求中读取文件
		f, err := c.FormFile("f1") //从请求中获取携带的参数一样
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			// 讲读取到的文件保存在本地(服务端本地)
			// dst := fmt.Sprintf("./%s", f.Filename)
			dst := path.Join("./", f.Filename)
			_ = c.SaveUploadedFile(f, dst)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	// 处理多个文件上传
	r.POST("/uploadMulti", func(c *gin.Context) {
		// 从请求中读取文件
		form, _ := c.MultipartForm()
		files := form.File["f2"]

		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("./%s_%d", file.Filename, index)
			_ = c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d file uploaded", len(files)),
		})
	})

	r.Run(":8080")
}
