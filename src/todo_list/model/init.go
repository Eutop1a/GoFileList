package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

func Database(connstring string) {
	db, err := gorm.Open("mysql", connstring)
	if err != nil {
		panic("Mysql数据库连接错误")
	}
	fmt.Println("Mysql数据库连接成功")
	// 打印日志
	db.LogMode(true)
	// 如果是发行版就不要输出日志
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	db.SingularTable(true)                       // 表名不加s
	db.DB().SetMaxIdleConns(20)                  // 设置连接池
	db.DB().SetMaxOpenConns(1000)                // 最大连接数
	db.DB().SetConnMaxLifetime(time.Second * 30) // 设置连接时间
	DB = db
	migration()
}
