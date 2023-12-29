package routes

import (
	"OJReconstitution/api"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession", store))
	// 跨域
	//r.Use(middleware.Core())

	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "success")
		})
		// 处理用户请求
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
		// 创建一个需要登录保护的路由组
		authed := v1.Group("/")
		{
			authed.POST("")
		}
	}
	return r
}
