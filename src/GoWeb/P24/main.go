package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User 1、定义模型
type User struct {
	// gorm.Model // 里面有	CreatedAt UpdatedAt DeletedAt
	ID     int64
	Name   string
	Age    int64
	Active bool
}

func main() {
	// 连接数据库
	db, err := gorm.Open("mysql",
		"root:123456@(127.0.0.1:3306)/db2?charset=utf8mb4&parseTime=True&loc-Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	//U1 := User{
	//	Name:   "q1mi",
	//	Age:    30,
	//	Active: true,
	//}
	//db.Create(&U1)
	//
	//u2 := User{
	//	Name:   "jinzhu2",
	//	Age:    20,
	//	Active: false,
	//}
	//db.Create(&u2)

	// 删除
	//
	//var u = User{}
	////u.ID = 1
	//u.Name = "hello"
	//db.Debug().Delete(&u) // 如果主键为空，会删除所有model的所有记录

	db.Debug().Where("name=?", "jinzhu2").Delete(User{})
	db.Debug().Delete(User{}, "age=?", 23)

	// model中有DeleteAt字段时，自动获得软删除
	//var u1 []User
	// 查得到
	// db.Debug().Unscoped().Where("name=?", "hello").Find(&u1)
	// 查不到
	//db.Debug().Where("name=?", "hello").Find(&u1)
	//
	//fmt.Println(u1)
	//
	//db.Debug().Unscoped().Where("name=?", "hello").Delete(User{})

}
