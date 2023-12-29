package services

import (
	"OnlineJudge/dao/mysql/users"
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

type Check struct {
	UserName string `form:"username"`
	NewScore int    `form:"newscore"`
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

func (check *Check) CheckUserName() (int, string) {
	// 判断用户名是否已存在
	status := users.ContainUser(check.UserName)
	if status {
		// 用户名已存在
		zap.L().Error(fmt.Sprintf("username %s already exists", check.UserName))
		return http.StatusForbidden, "UserExist"
	} else {
		// 用户名不存在
		return http.StatusOK, "notExist"
	}
}
