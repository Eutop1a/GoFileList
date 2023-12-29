package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	gorm.Model   // 内嵌gorm.Model
	Name         string
	Age          sql.NullInt64 `gorm:"column:user_age"` // 改列名
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`
	MemberNumber *string `gorm:"unique;not null"`
	Num          int     `gorm:"AUTO_INCREMENT"`
	Address      string  `gorm:"index:addr"` // address创建姓名为addr的索引
	IgnoreMe     int     `gorm:"-"`          // 忽略本字段
}

type Animal struct {
	// 默认用ID字段当主键，但是可以：
	// 用tag字段指定主键
	AnimalID int64 `gorm:"primary_key"`
	Name     string
	Age      int64
}

// TableName 唯一指定表明
func (Animal) TableName() string {
	return "q1mi"
}

func main() {
	// 给所有默认表名添加前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "prefix_" + defaultTableName
	}
	// 连接数据库
	db, err := gorm.Open("mysql",
		"root:123456@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc-Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.SingularTable(true) // 禁用复数

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Animal{})

	// 使用User结构体创建名叫xiaowangzi的表
	db.Table("xiaowangzi").CreateTable(&User{})

}
