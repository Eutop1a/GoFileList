package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
	"todo_list/pkg/utils"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := 200
		// var data interface{}
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claim, err := utils.ParseToken(token)
			if err != nil {
				code = 403 // 无权限，说明token是假的
			} else if time.Now().Unix() > claim.ExpiresAt {
				// token无效
				code = 401
			}
		}
		if code != 200 {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    "Token解析错误",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
