package controller

import (
	"OnlineJudge/pkg/jwt"
	"OnlineJudge/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

// Login 用户登录
// @Summary 用户登录
// @Description 用户使用用户名和密码进行登录，成功返回Token
// @Accept json
// @Produce json
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {object} models._ResponseToken "成功登录，返回Token"
// @Failure 403 {object} models._ResponseMsg "用户名不存在或密码错误"
// @Router /login [post]
func Login(c *gin.Context) {
	var userLogin services.UserService
	if err := c.ShouldBind(&userLogin); err == nil {
		statusCode, token := userLogin.Login()
		if statusCode == http.StatusOK {
			c.JSON(http.StatusOK, gin.H{
				"Token": token,
			})
		} else if statusCode == http.StatusForbidden {
			c.JSON(http.StatusForbidden, gin.H{
				"Msg": "username do not exist or wrong password",
			})
		}
	} else {
		zap.L().Error("SignUp with invalid username or password")
		c.JSON(http.StatusForbidden, gin.H{
			"Msg": "username do not exist or wrong password",
		})
	}

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
func Register(c *gin.Context) {

	var userRegister services.UserService
	if err := c.ShouldBind(&userRegister); err == nil {
		statusCode, token := userRegister.Register()
		if statusCode == http.StatusForbidden {
			c.JSON(http.StatusForbidden, gin.H{
				"Msg": "UserExist",
			})
		} else if statusCode == http.StatusInternalServerError {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Msg": "RegisterError",
				"Err": err.Error(),
			})
		} else if statusCode == http.StatusOK {
			c.JSON(http.StatusOK, gin.H{
				"Token":   token,
				"EndTime": time.Now().Add(jwt.TokenTime).String(),
			})
		}

	} else {
		zap.L().Error("Signup with invalid username or password")
		c.JSON(http.StatusForbidden, gin.H{
			"Msg": "Signup with invalid username or password",
		})
	}

}

func RegAdmin(c *gin.Context) {
	var regAdmin services.UserService
	// 从请求中获取执行操作的用户名
	tmp := c.Request.Form.Get("username")
	if tmp == "" {
		// Token 已超时
		zap.L().Error("TokenTimeOut")
		c.JSON(http.StatusForbidden, gin.H{"Msg": "TokenTimeout"})
		return
	} else {
		regAdmin.UserName = tmp
	}

	if err := c.ShouldBind(&regAdmin); err == nil {
		statusCode, Msg := regAdmin.RegAdmin()
		// 注册成功
		if statusCode == http.StatusCreated {
			c.JSON(http.StatusCreated, gin.H{
				"Msg": Msg,
			})
			// author不对
		} else if statusCode == http.StatusForbidden {
			c.JSON(http.StatusForbidden, gin.H{
				"Msg": Msg,
			})
			// 内部数据库错误
		} else if statusCode == http.StatusInternalServerError {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Msg": Msg,
			})
		} else if statusCode == http.StatusConflict {
			c.JSON(http.StatusConflict, gin.H{
				"Msg": Msg,
			})
		}

	} else {
		zap.L().Error("invalid parameter")
		c.JSON(http.StatusForbidden, gin.H{
			"Msg": "username do not exist or wrong password",
		})
	}
}
