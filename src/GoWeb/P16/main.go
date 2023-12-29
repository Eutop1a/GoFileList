package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 重定向

func main() {
	r := gin.Default()

	r.GET("/index", func(c *gin.Context) {
		// 跳转到百度
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	r.GET("/a", func(c *gin.Context) {
		// 跳转到 /b 对应的路由处理函数
		c.Request.URL.Path = "/b" // 把请求的URL修改
		r.HandleContext(c)        // 继续后续处理
	})

	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "b",
		})
	})

	r.Run(":8080")
}
