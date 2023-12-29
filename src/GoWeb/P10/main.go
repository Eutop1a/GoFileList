package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		// 方法1: 使用map
		//data := map[string]interface{}{
		//	"name":    "wcl",
		//	"message": "Hello",
		//	"age":     19,
		//}
		// gin内置的
		data := gin.H{
			"name":    "wcl",
			"message": "Hello",
			"age":     20,
		}
		c.JSON(http.StatusOK, data)
	})

	// 方法2：结构体 灵活使用tag实现定制化操作
	type msg struct {
		Name    string `json:"name"`
		Message string
		Age     int
	}
	r.GET("/another_json", func(c *gin.Context) {
		data := msg{
			Name:    "wcl",
			Message: "你好",
			Age:     20,
		}
		c.JSON(http.StatusOK, data) // json序列化(通过反射实现数据的序列化)
	})
	r.Run(":8080")
}
