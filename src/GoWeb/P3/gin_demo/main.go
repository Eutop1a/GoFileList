package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello Golang",
	})
}

func main() {
	// 返回默认的路由引擎
	r := gin.Default()

	//指定用户使用GET请求访问/hello时，执行sayHello
	r.GET("/hello", sayHello)
	/*
		r.GET("/book", sayHello)
		r.POST("/book", sayHello)
		r.PUT("/book", sayHello)
		r.DELETE("/book", sayHello)
	*/
	// RESTful API
	r.GET("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})
	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})
	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})
	// 启动服务
	r.Run(":9090")
}
