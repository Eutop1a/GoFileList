package main

// ORM
// Object Relational Mapping
// 对象		关系		  映射
// go结构体  关系型数据库

// 优点：提高开发效率
// 缺点：牺牲执行性能，灵活性，弱化SQL能力

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo --> 数据表
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	// 连接数据库
	db, err := gorm.Open("mysql",
		"root:123456@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc-Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 创建表 自动迁移（把结构体和数据表进行对应）
	db.AutoMigrate(&UserInfo{})

	// 创建数据行
	//u1 := UserInfo{
	//	ID:     0,
	//	Name:   "me",
	//	Gender: "man",
	//	Hobby:  "篮球",
	//}
	//db.Create(&u1)

	// 查询
	var u UserInfo
	db.First(&u) // 查询表中第一条数据保存到u中
	fmt.Printf("u:%#v\n", u)

	// 更新
	db.Model(&u).Update("hobby", "双色球")

	// 删除
	db.Delete(&u)
}
