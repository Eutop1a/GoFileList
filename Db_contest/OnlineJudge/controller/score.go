package controller

import (
	"OnlineJudge/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// AddScore 增加用户分数
// @Summary 增加用户分数
// @Description 通过用户名和新分数的表单数据，增加用户的总分数，并返回新的总分数
// @Accept json
// @Produce json
// @Param username formData string true "用户名"
// @Param newscore formData string true "要增加的分数"
// @Success 200 {object} models._ResponseAddScore "增加分数成功，返回消息和新的总分数"
// @Failure 403 {string} models._ResponseMsg  "Token 已超时"
// @Failure 500 {object} models._ResponseError "数据库查询或保存出错"
// @Router /changeScore [post]
func AddScore(c *gin.Context) {
	var addScore services.ChangeScore
	if err := c.ShouldBind(&addScore); err == nil {
		tmp := c.Request.Form.Get("username")
		if tmp == "" {
			zap.L().Error("TokenTimeout")
			c.JSON(http.StatusForbidden, gin.H{"Msg": "TokenTimeout"})
			return
		} else {
			addScore.UserName = tmp
		}
		statusCode, errMsg, ret := addScore.AddScore()
		if statusCode == http.StatusOK {
			c.JSON(http.StatusOK, gin.H{
				"msg":       "change score success",
				"new_score": ret,
			})
		} else if statusCode == http.StatusInternalServerError {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Err": errMsg,
			})
		} else if statusCode == http.StatusForbidden {
			c.JSON(http.StatusForbidden, gin.H{
				"Err": errMsg,
			})
		}
	} else {
		zap.L().Error("SignUp with invalid username or password")
		c.JSON(http.StatusForbidden, gin.H{
			"Msg": "username do not exist or wrong password",
		})
	}
}

// SortByScore 根据用户分数降序排序
// @Summary 根据用户分数降序排序
// @Description 获取所有用户并按分数降序排序
// @Accept json
// @Produce json
// @Param username formData string true "执行操作的用户名"
// @Success 200 {object} models._ResponseSort "成功获取并排序用户列表"
// @Failure 403 {object} models._ResponseMsg "Token 已超时"
// @Failure 500 {object} models._ResponseError "数据库查询错误"
// @Router /api/sort [post]
func SortByScore(c *gin.Context) {
	var scoreSort services.ChangeScore
	if err := c.ShouldBind(&scoreSort); err == nil {
		tmp := c.Request.Form.Get("username")
		if tmp == "" {
			zap.L().Error("TokenTimeout")
			c.JSON(http.StatusForbidden, gin.H{"Msg": "TokenTimeout"})
			return
		} else {
			scoreSort.UserName = tmp
		}

		statusCode, errMsg, data := scoreSort.SortByScore()

		if statusCode == http.StatusOK {

			c.JSON(http.StatusOK, gin.H{
				"Msg":  "Get sort by score",
				"Data": data,
			})
			return

		} else if statusCode == http.StatusInternalServerError {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Err": errMsg,
			})
			return
		}

		return
	} else {
		zap.L().Error("SignUp with invalid username or password")
		c.JSON(http.StatusForbidden, gin.H{
			"Msg": "username do not exist or wrong password",
		})
		return
	}
}
