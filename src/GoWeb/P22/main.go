package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User 1、定义模型
type User struct {
	gorm.Model // 里面有	CreatedAt UpdatedAt DeletedAt
	Name       string
	Age        int64
}

func main() {
	// 连接数据库
	db, err := gorm.Open("mysql",
		"root:123456@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc-Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 2、把模型与数据库中的表对应起来

	db.AutoMigrate(&User{})

	// 3、创建

	//u1 := User{
	//	Name: "q1mi",
	//	Age:  18,
	//}
	//db.Create(&u1)
	//u2 := User{
	//	Name: "jinzhu",
	//	Age:  20,
	//}
	//db.Create(&u2)

	// 4、查询
	//var user User // 声明模型结构体变量user
	user := new(User) // new返回的是基本变量类型的指针
	// new和make的区别
	db.Debug().First(user)
	// SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL ORDER BY `users`.`id` ASC LIMIT 1
	fmt.Printf("user:%v\n", user)
	var users []User
	db.Debug().Find(&users)
	// SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL
	fmt.Printf("users:%v\n", users)

	// FirstOrInit
	var U User
	//db.Attrs(User{Age: 19}).FirstOrInit(&U, User{
	//	Name: "xwz",
	//})
	db.Attrs(User{Age: 19}).FirstOrInit(&U, User{
		Name: "jinzhu",
	})
	fmt.Printf("users:%v\n", U)

}
