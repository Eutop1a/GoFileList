package main

import (
	"P26/dao"
	"P26/models"
	"P26/routers"
)

func main() {
	// 创建数据库
	// sql : create database bubble;
	// 连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close() //程序退出就关闭数据库连接

	//dao.InitModel()

	dao.DB.AutoMigrate(&models.Todo{})

	r := routers.SetupRouters()

	r.Run(":8080")
}
