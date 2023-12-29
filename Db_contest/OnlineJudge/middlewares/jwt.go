package middlewares

import (
	"OnlineJudge/pkg/jwt"
	"github.com/gin-gonic/gin"
)

// AuthorizationMiddleWare 认证中间件
// @Summary 认证中间件，验证请求中的授权令牌
// @Description 从请求头中提取授权令牌，验证其有效性，如果有效则将用户名添加到请求表单中
// @Produce json
// @Param Authorization header string true "授权令牌，形如 'Bearer {token}'"
// @Success 200 {object} models._ResponseMsg "成功验证授权令牌"
// @Failure 401 {object} models._ResponseMsg "授权令牌无效或缺失"
// // @Router /api [middleware]
func AuthorizationMiddleWare(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	var username string
	var succ bool
	if token[:6] == "Bearer" {
		token = token[7:]
	}
	if username, succ = jwt.JudgeAccessToken(token); succ {

		c.Request.ParseForm()
		c.Request.Form.Set("username", username)

		//// 获取修改后的username值
		//newUsername := c.Request.Form.Get("username")
		//fmt.Println("New username:", newUsername)
		//c.Params["username"] = username
		c.Next()
	} else {
		//delete(c.Params, "username")
		c.Request.ParseForm()
		c.Request.Form.Del("username")
		c.Next()
	}
}
