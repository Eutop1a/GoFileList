package users

import (
	"OnlineJudge/dao/mysql"
	"OnlineJudge/models"
)

func ContainUser(username string) bool {

	var user models.User

	mysql.Db.First(&user, "username = ?", username)

	if user.Username != username {
		//fmt.Println(user.Username, "!=", username)
		return false
	}
	//fmt.Println(user.Username, "!=", username)
	return true
}
