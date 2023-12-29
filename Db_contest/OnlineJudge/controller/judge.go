package controller

import (
	"OnlineJudge/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// JudgeCode 代码评测
// @Summary 代码评测
// @Description 通过用户身份、问题ID和代码内容，进行代码评测，并返回评测结果
// @Accept json
// @Produce json
// @Param username formData string true "执行操作的用户名"
// @Param problem formData integer true "要评测的问题ID"
// @Param code formData string true "要评测的代码内容"
// @Success 200 {object} models._ResponseMsg "成功进行代码评测"
// @Failure 403 {object} models._ResponseMsg "Token 已超时"
// @Router /judge [post]
func JudgeCode(c *gin.Context) {
	var judge services.Judge
	// 从请求中获取执行操作的用户名
	tmp := c.Request.Form.Get("username")

	if tmp == "" {
		// Token 已超时
		zap.L().Error("TokenTimeOut")
		c.JSON(http.StatusForbidden, gin.H{"Msg": "TokenTimeout"})
		return
	} else {
		judge.UserName = tmp
	}
	// 获取代码
	if err := c.ShouldBind(&judge); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Msg": "Request error",
		})
	}

	statusCode, Msg := judge.JudgeCode()
	if statusCode == http.StatusOK {
		c.JSON(http.StatusOK, gin.H{
			"Msg": Msg,
		})
	} else if statusCode == http.StatusInternalServerError {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Msg": Msg,
		})
	}
}
