package routes

import (
	"OnlineJudge/api"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"Msg": "success",
		})
	})
	// 用户的登录和注册操作
	user := r.Group("user")
	{
		user.POST("/login", api.UserRegister)
		user.POST("/register", api.UserLogin)
	}
	return r
}
