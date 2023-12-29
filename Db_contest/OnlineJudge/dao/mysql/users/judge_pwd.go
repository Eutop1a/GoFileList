package users

import (
	"OnlineJudge/dao/mysql"
	"OnlineJudge/models"
)

func JudgePWD(username string, pwd string) bool {

	var user models.User
	mysql.Db.First(&user, "username = ?", username)

	if user.Username != username {
		return false
	}

	return user.Password == pwd

}
