package mysql

import (
	"bluebell/setting"
	"fmt"

	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init(cfg *setting.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
		//viper.GetString("mysql.user"),
		//viper.GetString("mysql.password"),
		//viper.GetString("mysql.host"),
		//viper.GetInt("mysql.port"),
		//viper.GetString("mysql.dbname"),
	)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed, err: %v\n", zap.Error(err))
		return
	}
	//db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
	//db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	return
}

func Close() {
	_ = db.Close()
}
