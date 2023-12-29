package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 路由与路由器组

func main() {
	r := gin.Default()
	// 访问/index的GET请求会走这一条处理逻辑
	// 路由
	// r.HEAD()
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})
	// 新建
	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})
	// 更新
	r.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})
	// 删除
	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})
	// any:请求方法大集合
	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK, gin.H{
				"method": "Get",
			})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{
				"method": "Post:",
			})
			// ...
		}
	})

	// NoRoute 不存在的的路由走这里
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"liwenzhou.com": "nihao",
		})
	})
	// 视频的首页和详情页
	//r.GET("/video/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "/video/index",
	//	})
	//})

	// 路由组: 提取公用的前缀
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/video/index",
			})
		})
		videoGroup.GET("/xx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/video/xx",
			})
		})
		videoGroup.GET("/oo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/video/oo",
			})
		})
	}

	// 商城的首页和详情页
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/shopGroup/index",
			})
		})
		shopGroup.GET("/xx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/shopGroup/xx",
			})
		})
		shopGroup.GET("/oo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/shopGroup/oo",
			})
		})

	}
	r.Run(":8080")
}
