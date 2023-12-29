package users

import (
	"OnlineJudge/dao/mysql"
	"OnlineJudge/models"
	"crypto/md5"
	"fmt"
)

const Title = "Eutop1a"

func NewUser(username, password string) error {

	pwd := fmt.Sprintf("%x", md5.Sum([]byte(password+Title)))

	mysql.Db.Create(&models.User{
		Username: username,
		Password: pwd,
	})

	return nil
}
