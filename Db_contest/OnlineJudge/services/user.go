package services

import (
	"OnlineJudge/dao/mysql/users"
	"OnlineJudge/pkg/jwt"
	"crypto/md5"
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

const Title = "Eutop1a"

type UserService struct {
	UserName string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Author   string `form:"author" json:"author"`
}

// Register 用户注册
// @Summary 用户注册
// @Description 用户使用用户名、密码和昵称进行注册，成功返回Token
// @Accept json
// @Produce json
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Param nickname formData string true "昵称"
// @Success 200 {object} models._ResponseReg "成功注册，返回Token和Token过期时间"
// @Failure 403 {object} models._ResponseRegErr "用户已存在或注册失败"
// @Router /register [post]
func (service *UserService) Register() (int, string) {
	username := service.UserName
	password := service.Password

	if users.ContainUser(service.UserName) {
		zap.L().Error(fmt.Sprintf("username %s already exists", username))
		return http.StatusForbidden, ""
	}

	err := users.NewUser(username, password)
	if err != nil {
		zap.L().Error("RegisterError", zap.Error(err))
		return http.StatusInternalServerError, ""
	} else {
		return http.StatusOK, jwt.GetToken(username)
	}
}

func (service *UserService) Login() (int, string) {
	username := service.UserName
	password := service.Password
	if users.JudgePWD(username, fmt.Sprintf("%x", md5.Sum([]byte(password+Title)))) {
		return http.StatusOK, jwt.GetToken(username)
	} else {
		zap.L().Error("SignUp with invalid username or password")
		return http.StatusForbidden, ""
	}
}

func (service *UserService) RegAdmin() (int, string) {
	if service.Author != Title {
		zap.L().Error("Title error, permission deny")
		return http.StatusForbidden, "Title error, permission deny"
	} else {
		ok := users.RegAdmin(service.UserName)
		if ok == 2 {
			zap.L().Error("RegAdmin error, internal mysql error")
			return http.StatusInternalServerError, "Add admin error"
		} else if ok == 1 {
			zap.L().Error(fmt.Sprintf("User %s aleardy is admin", service.UserName))
			return http.StatusConflict, fmt.Sprintf("User %s aleardy is admin", service.UserName)
		}
		return http.StatusCreated, "Reg admin success"
	}

}
