package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User 1、定义模型
type User struct {
	ID int64
	// 有默认值的情况下传入空的做法：
	// 1、Name: new(string), (方法是把定义中的string类型改为*string)
	// 2、改类型为sql.NullString
	// Name *string `gorm:"default:'wcl'"`
	// Name *string `gorm:"default:'wcl'"`
	Name sql.NullString `gorm:"default:'wcl'"`
	Age  int64
}

func main() {
	db, err := gorm.Open("mysql",
		"root:123456@(127.0.0.1:3306)/db2?charset=utf8mb4&parseTime=True&loc-Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// 2、把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})

	// 3、创建
	u := User{
		// 有默认值的情况下传入空的做法：
		// 1、Name: new(string), (方法是把定义中的string类型改为*string)
		// 2、改类型为sql.NullString
		Name: sql.NullString{
			String: "",
			Valid:  true,
		},
		Age: 18,
	} // 在代码层面创建一个User对象
	fmt.Println(db.NewRecord(&u)) // 判断主键是否为空 true
	db.Debug().Create(&u)         // 在数据库中创建了一条记录
	fmt.Println(db.NewRecord(&u)) // 判断主键是否为空 false
}
