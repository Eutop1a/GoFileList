package main

// 中间件
// 适合处理一些公共的业务逻辑（登录认证，权限校验，数据分页，记录日志，耗时统计）

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// handleFunc
func indexHandler(c *gin.Context) {
	fmt.Println("index")
	name, ok := c.Get("name") // 从上下文中取值（跨中间件存取值）
	if !ok {
		name = "匿名用户"
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": name,
	})
}

// 定义一个中间件 m1 统计请求处理函数耗时
func m1(c *gin.Context) {
	fmt.Println("m1 in...")
	// 计时
	start := time.Now()
	go func(c *gin.Context) {

	}(c.Copy()) // 在func中只能使用c的拷贝，不能直接使用c
	c.Next()    // 调用后续的处理函数
	// c.Abort() // 阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("m1 out")
}

// 定义一个中间件 m2 统计请求处理函数耗时
func m2(c *gin.Context) {
	fmt.Println("m2 in...")
	c.Set("name", "q1mi") // 在上下文c中设置值
	//c.Next()
	//c.Abort()
	//return // 不执行后面的print
	fmt.Println("m2 out...")
}

// 登录认证中间件
func authMiddleware(doCheck bool) gin.HandlerFunc {
	// 连接数据库
	// 或者做一些其他准备工作
	return func(c *gin.Context) {
		if doCheck {
			// 是否登录
			// if 是登录
			c.Next()
			// else
			// c.Abort()
		} else {
			c.Next()
		}
	}
}

func main() {
	// 默认使用了Logger()和Recover()y两个中间件
	// r := gin.Default()
	r := gin.New()
	// 全局注册中间件函数
	r.Use(m1, m2, authMiddleware(true))

	r.GET("/index", indexHandler)

	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "shop",
		})
	})

	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "user",
		})
	})
	// 路由组注册中间件方法1
	xxGroup := r.Group("/xx", authMiddleware(true))
	{
		xxGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "xxGroup"})
		})
	}
	xxGroup2 := r.Group("/xx2")
	xxGroup2.Use(authMiddleware(true))
	{
		xxGroup2.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "xxGroup"})
		})
	}
	go func() {}()
	r.Run(":8080")
}
