package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取请求的path(URL)参数，返回的都是string类型

func main() {
	r := gin.Default()

	r.GET("/user/:name/:age", func(c *gin.Context) {
		// 获取路径参数
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.GET("/blog/:year/:month", func(c *gin.Context) {
		// 获取路径参数
		year := c.Param("year")
		month := c.Param("month")
		c.JSON(http.StatusOK, gin.H{
			"year":  year,
			"month": month,
		})
	})

	r.Run(":8080")
}
