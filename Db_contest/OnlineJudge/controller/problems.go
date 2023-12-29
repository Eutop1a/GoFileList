package controller

import (
	"OnlineJudge/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

// GetList 获取问题列表
// @Summary 获取问题列表
// @Description 获取所有问题的列表
// @Accept json
// @Produce json
// @Param username formData string true "要获取问题列表的用户名"
// @Success 200 {object} models._ResponseProblems "成功获取问题列表"
// @Failure 403 {object} models._ResponseMsg "Token 已超时"
// @Failure 500 {object} models._ResponseMsg "获取问题列表出错"
// @Router /list [post]
func GetList(c *gin.Context) {
	var problem services.Problem
	// 从请求中获取执行操作的用户名
	tmp := c.Request.Form.Get("username")

	if tmp == "" {
		// Token 已超时
		zap.L().Error("TokenTimeOut")
		c.JSON(http.StatusForbidden, gin.H{
			"Msg": "TokenTimeOut",
		})
		return
	} else {
		problem.UserName = tmp
	}

	statusCode, Msg, data := problem.GetList()
	if statusCode == http.StatusInternalServerError {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Msg": Msg,
		})

	} else if statusCode == http.StatusOK {
		c.JSON(http.StatusOK, gin.H{
			"Msg":  "Get problem success",
			"Data": data,
		})
	}

}

// AddProblem 添加问题
// @Summary 添加问题
// @Description 通过管理员身份，使用表单数据添加新的问题，并返回操作结果
// @Accept json
// @Produce json
// @Param username formData string true "执行操作的管理员用户名"
// @Param id formData string true "问题ID"
// @Param title formData string true "问题标题"
// @Param lore formData string true "问题描述"
// @Param input formData string true "问题标准输入"
// @Param output formData string true "问题标准输出"
// @Param tips formData string true "问题提示"
// @Success 200 {object} models._ResponseAddProblems "成功添加问题"
// @Failure 403 {object} models._ResponseMsg "Token 已超时或用户非管理员"
// @Failure 403 {object} models._ResponseMsg "解析表单数据出错"
// @Router /problem/add [post]
func AddProblem(c *gin.Context) {
	var problem services.Problem
	// 从请求中获取执行操作的用户名
	tmp := c.Request.Form.Get("username")

	if tmp == "" {
		// Token 已超时
		zap.L().Error("TokenTimeOut")
		c.JSON(http.StatusForbidden, gin.H{
			"Msg": "TokenTimeOut",
		})
		return
	} else {
		problem.UserName = tmp
	}

	// 处理文件数据（如果有）

	// 解析请求中的 multipart/form-data 数据
	err := c.Request.ParseMultipartForm(32 << 20) // 32 << 20 是最大允许的内存大小
	if err != nil {
		zap.L().Error("ParseMultipartForm", zap.Error(err))
		c.JSON(http.StatusForbidden, gin.H{
			"Msg": "ParseMultipartForm",
		})
		return
	}

	// 获取表单字段
	problem.Id = c.Request.FormValue("id")
	problem.Title = c.Request.FormValue("title")
	problem.Lore = c.Request.FormValue("lore")
	problem.Input = c.Request.FormValue("input")
	problem.Output = c.Request.FormValue("output")
	problem.Tips = c.Request.FormValue("tips")
	problem.Score, _ = strconv.Atoi(c.Request.FormValue("score"))

	statusCode, Msg := problem.AddProblem()
	if statusCode == http.StatusForbidden {
		c.JSON(http.StatusForbidden, gin.H{
			"Msg": Msg,
		})
	} else if statusCode == http.StatusOK {
		// 返回成功响应
		c.JSON(http.StatusOK, gin.H{
			"Msg": Msg,
		})
	}
}

// GetProblem 获取问题详情
// @Summary 获取问题详情
// @Description 通过问题ID和管理员身份，获取指定问题的详细信息
// @Accept json
// @Produce json
// @Param _ formData string true "执行操作的管理员用户名"
// @Param id path string true "要获取的问题ID"
// @Success 200 {object} models._ResponseQuestionDetail  "成功获取问题详情"
// @Failure 403 {object} models._ResponseMsg "Token 已超时"
// @Failure 403 {object} models._ResponseMsg "获取问题详情出错"
// @Router /problem/:id [post]
func GetProblem(c *gin.Context) {
	var problem services.Problem
	// 从请求中获取执行操作的用户名
	tmp := c.Request.Form.Get("username")

	if tmp == "" {
		// Token 已超时
		zap.L().Error("TokenTimeOut")
		c.JSON(http.StatusForbidden, gin.H{
			"Msg": "TokenTimeOut",
		})
		return
	} else {
		problem.UserName = tmp
	}
	statusCode, Msg, data := problem.GetProblem(c.Param("id"))
	if statusCode == http.StatusForbidden {
		c.JSON(http.StatusForbidden, gin.H{
			"Msg": Msg,
		})
	} else if statusCode == http.StatusOK {
		// 返回成功响应
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
