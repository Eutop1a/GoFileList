package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"todo_list/api"
	"todo_list/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // 开启swag

	r.Use(sessions.Sessions("mysession", store))
	// 跨域
	r.Use(middleware.Cors())
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		// 用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
		authed := v1.Group("/") //需要登录保护
		authed.Use(middleware.JWT())
		{
			authed.POST("task", api.CreateTask)
			authed.GET("task/:id", api.ShowTask)
			authed.GET("tasks/", api.ListTask)
			authed.PUT("task/:id", api.UpdateTask)
			authed.POST("search", api.SearchTask)
			authed.DELETE("task/:id", api.DeleteTask)
		}
	}
	return r
}
