package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	ratelimit2 "github.com/juju/ratelimit" // 令牌
	ratelimit1 "go.uber.org/ratelimit"     // 漏桶
)

func pingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func heiHandler(c *gin.Context) {
	c.String(http.StatusOK, "ha")
}

// 基于漏桶的限流中间件1
func rateLimit1() func(ctx *gin.Context) {
	// 生成一个限流器
	rl := ratelimit1.New(100)
	return func(c *gin.Context) {
		if rl.Take().Sub(time.Now()) > 0 {
			// time.Sleep(rl.Take().Sub(time.Now())) // 需要等待这么长时间
			c.String(http.StatusOK, "rate limit...")
			c.Abort()
			return
		}
		c.Next()
	}
}

// 基于令牌桶的限流中间件2
func rateLimit2(fileInterval time.Duration, cap int64) func(ctx *gin.Context) {
	rl := ratelimit2.NewBucket(fileInterval, cap)
	return func(c *gin.Context) {
		//rl.Take()          // 可以欠账
		//rl.TakeAvailable() 不可以欠账(有令牌才会取出)
		if rl.TakeAvailable(1) == 1 { // 此次没有取到令牌
			c.Next()
			return
		}
		c.String(http.StatusOK, "rate limit...")
		c.Abort()
	}
}

func main() {
	r := gin.Default()

	r.GET("/ping", rateLimit1(), pingHandler)
	r.GET("/hei", rateLimit2(time.Second, 3), heiHandler)

	r.Run()

}
