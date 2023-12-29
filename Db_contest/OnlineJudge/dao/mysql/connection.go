package mysql

import (
	"OnlineJudge/models"
	"OnlineJudge/setting"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Init(cfg *setting.MySQLConfig) error {

	//dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?parseTime=true", username, password, protocal, address, port, database)
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Protocal,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Error("connect DB failed, err: %v\n", zap.Error(err))
		return err
	}
	Db.AutoMigrate(&models.User{})
	//Db.AutoMigrate(&models.UserSolved{})
	Db.AutoMigrate(&models.ProblemList{})
	Db.AutoMigrate(&models.Problems{})
	Db.AutoMigrate(&models.Admins{})

	return err
}
