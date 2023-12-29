package users

import (
	"OnlineJudge/dao/mysql"
	"OnlineJudge/models"
)

func JudgeAdmin(username string) bool {

	var admin models.Admins

	mysql.Db.First(&admin, "username = ?", username)

	return admin.Username == username

}

func RegAdmin(username string) int {
	var existingAdmin models.Admins

	// 检查用户名是否已经存在
	err := mysql.Db.First(&existingAdmin, "username = ?", username).Error
	if err == nil {
		// 用户名已经存在，不进行创建
		return 1
	}

	// 创建新的管理员记录
	newAdmin := models.Admins{
		Username: username,
	}

	err = nil
	err = mysql.Db.Create(&newAdmin).Error
	if err != nil {
		// 内部错误
		return 2
	}
	// 返回true表示成功创建管理员
	return 0
}
