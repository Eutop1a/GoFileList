package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

func Database(conn string) {
	Db, err := gorm.Open("mysql", conn)
	if err != nil {
		panic("数据库连接错误")
	}
	fmt.Println("Mysql连接成功")
	// 打印日志
	Db.LogMode(true)
	// 如果是发行版就不要输出日志
	if gin.Mode() == "release" {
		Db.LogMode(false)
	}
	Db.SingularTable(true)
	Db.DB().SetMaxIdleConns(200)
	Db.DB().SetMaxOpenConns(1000)
	Db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = Db
}
