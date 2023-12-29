package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User 1、定义模型
type User struct {
	gorm.Model // 里面有	CreatedAt UpdatedAt DeletedAt
	Name       string
	Age        int64
	Active     bool
}

func main() {
	// 连接数据库
	db, err := gorm.Open("mysql",
		"root:123456@(127.0.0.1:3306)/db2?charset=utf8mb4&parseTime=True&loc-Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//
	//db.AutoMigrate(&User{})
	//
	//u1 := User{
	//	Name:   "q1mi",
	//	Age:    18,
	//	Active: true,
	//}
	//db.Create(&u1)
	//u2 := User{
	//	Name:   "jinzhu",
	//	Age:    20,
	//	Active: false,
	//}
	//db.Create(&u2)

	// select
	var use User
	db.First(&use)

	// update
	use.Name = "七米"
	use.Age = 21

	db.Debug().Save(&use) // 默认修改所有字段

	//db.Debug().Model(&use).Update("name", "小王子")
	//
	//m1 := map[string]interface{}{
	//	"name":   "wcl",
	//	"age":    20,
	//	"active": true,
	//}
	//
	//db.Debug().Model(&use).Updates(m1)               // m1列出来的所有字段都更新
	//db.Debug().Model(&use).Select("age").Update(m1)  // 只更新age字段
	//db.Debug().Model(&use).Omit("active").Update(m1) // 排除m1中active字段更新其他的字段
	//
	//db.Debug().Model(&use).UpdateColumn("age", 30) // 排除m1中active字段更新其他的字段
	//
	//rowsNum := db.Model(User{}).Updates(User{
	//	Name: "hello",
	//	Age:  18,
	//}).RowsAffected
	//fmt.Println(rowsNum)
	//
	//

	// 让users表中所有用户年龄+2
	db.Debug().Model(&User{}).
		Update("age", gorm.Expr("age+?", 2))

}
