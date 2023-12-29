package configs

import (
	"OnlineJudge/db"
	"fmt"
	"gopkg.in/ini.v1"
	"log"
)

var (
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func Init() {
	file, err := ini.Load("./configs/config.ini")
	if err != nil {
		log.Fatalf("配置文件读取错误，请检查文件路径")
	}
	LoadServer(file)
	LoadMysql(file)
	// path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true"}, "")
	path := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true", DbUser, DbPassWord, DbHost, DbPort, DbName)
	db.Connection(path)
	//model.Database(path)
	//cache.Redis()
}

func LoadServer(file *ini.File) {
	// 读取ini中service下面的内容
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadMysql(file *ini.File) {
	// 读取ini中mysql下面的内容
	Db = file.Section("database1").Key("driver").String()
	DbHost = file.Section("database1").Key("host").String()
	DbPort = file.Section("database1").Key("port").String()
	DbUser = file.Section("database1").Key("username").String()
	DbPassWord = file.Section("database1").Key("password").String()
	DbName = file.Section("database1").Key("database").String()

}
