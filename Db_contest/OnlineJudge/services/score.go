package services

import (
	"OnlineJudge/dao/mysql"
	"OnlineJudge/models"
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

type ChangeScore struct {
	UserName  string `form:"username"`
	ProblemId string `form:"problem"`
}

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

func (score *ChangeScore) AddScore() (int, string, int) {
	// 从请求中获取用户名和新分数
	username := score.UserName
	problemId := score.ProblemId

	// 通过用户名查询用户信息
	var user models.User
	err := mysql.Db.Where("username = ?", username).First(&user).Error
	if err != nil {
		// 数据库查询出错
		zap.L().Error(fmt.Sprintf("database select %s error", username), zap.Error(err))
		return http.StatusInternalServerError, err.Error(), 0
	}

	// 获取题目分数
	var problem models.ProblemList
	err = mysql.Db.Where("problem_id = ?", problemId).First(&problem).Error
	if err != nil {
		// 数据库查询出错
		zap.L().Error(fmt.Sprintf("database get problem id: %s error", problemId), zap.Error(err))
		return http.StatusInternalServerError, err.Error(), 0
	}

	// 解析新分数并更新用户总分数

	user.Score += problem.ProblemScore

	// 保存更新后的用户信息
	err = mysql.Db.Save(&user).Error
	if err != nil {
		// 数据库保存出错
		zap.L().Error(fmt.Sprintf("database save %s error", username), zap.Error(err))

		return http.StatusInternalServerError, err.Error(), 0
	}

	// 返回成功响应
	return http.StatusOK, "", user.Score

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
func (score *ChangeScore) SortByScore() (int, string, []models.User) {
	// 从请求中获取执行操作的用户名
	//username := score.UserName

	// 查询并按分数降序排序用户列表
	var users []models.User
	err := mysql.Db.Order("score desc").Find(&users).Error
	if err != nil {
		zap.L().Error("database order error", zap.Error(err))
		return http.StatusInternalServerError, err.Error(), nil
	}

	return http.StatusOK, "", users
}
