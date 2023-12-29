package controller

import (
	"OnlineJudge/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// CheckToken 检查令牌有效性
// @Summary 检查令牌有效性
// @Description 通过表单数据中的用户名检查令牌的有效性，返回检查结果
// @Accept json
// @Produce json
// @Param username formData string true "要检查令牌的用户名"
// @Success 200 {object} models._ResponseUsername "令牌有效，返回用户名"
// @Failure 403 {object} models._ResponseMsg "无效令牌"
// @Router /open [get]

func CheckToken(c *gin.Context) {
	// 从请求中获取要检查令牌的用户名
	username := c.Request.Form.Get("username")

	if username == "" {
		// 无效令牌
		zap.L().Error("Invalid token")
		c.JSON(http.StatusForbidden, gin.H{
			"Msg": "Invalid token",
		})
	} else {
		// 令牌有效，返回用户名
		c.JSON(http.StatusOK, gin.H{
			"username": username,
		})
	}
}

// Check 检查用户名是否存在
// @Summary 检查用户名是否存在
// @Description 通过用户名的表单数据，检查该用户名是否已存在
// @Accept json
// @Produce json
// @Param username formData string true "要检查的用户名"
// @Success 200 {object} models._ResponseMsg "用户名不存在"
// @Failure 403 {object} models._ResponseMsg "Token 已超时或用户名已存在"
// @Router /check [post]

func CheckUserName(c *gin.Context) {
	var check services.Check
	// 从请求中获取要检查的用户名
	if err := c.ShouldBind(&check); err == nil {
		statusCode, Msg := check.CheckUserName()
		// 判断用户名是否已存在

		if statusCode == http.StatusForbidden {
			// 用户名已存在
			c.JSON(http.StatusForbidden, gin.H{
				"Msg": Msg,
			})
		} else {
			// 用户名不存在
			c.JSON(http.StatusOK, gin.H{
				"Msg": Msg,
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Msg": "request error",
		})
	}

}
