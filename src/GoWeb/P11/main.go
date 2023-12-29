package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// query string
	// GET请求 URL ? 后面的是query string参数
	// key=value格式，多个key-value用&连接
	// eq: /web?query=wcl&age=20

	r.GET("/web", func(c *gin.Context) {
		// 获取浏览器发请求携带的query String 参数
		name := c.Query("query") // 通过Query获取前端请求中携带的query String 参数
		age := c.Query("age")
		// name := c.DefaultQuery("query", "somebody") //取不到就用指定的默认值
		// name, ok := c.GetQuery("query") // (value, true)   ("", false)
		// if !ok {
		// 	name = "somebody"
		// }
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.Run(":8080")
}
