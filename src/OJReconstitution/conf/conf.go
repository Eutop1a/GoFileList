package conf

import (
	"OJReconstitution/model"
	"fmt"
	"gopkg.in/ini.v1"
	"strings"
)

var (
	AppMode    string
	HttpPort   string
	DB         string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassWord string
	DBName     string
)

func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径")
	}

	LoadServer(file)
	LoadMysql(file)
	path := strings.Join([]string{DBUser, ":", DBPassWord,
		"@tcp(", DBHost, ":", DBPort, ")/", DBName, "?charset=utf8mb4&parseTime=true"}, "")
	model.Database(path)
}

func LoadServer(file *ini.File) {
	// 读取ini中service下面的内容
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadMysql(file *ini.File) {
	// 读取ini中mysql下面的内容
	DB = file.Section("mysql").Key("DB").String()
	DBHost = file.Section("mysql").Key("DBHost").String()
	DBPort = file.Section("mysql").Key("DBPort").String()
	DBUser = file.Section("mysql").Key("DBUser").String()
	DBPassWord = file.Section("mysql").Key("DBPassWord").String()
	DBName = file.Section("mysql").Key("DBName").String()
}
