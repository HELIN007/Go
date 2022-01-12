package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

var (
	AppMode     string
	HttpPort    string
	JwtKey      string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassWord  string
	DBName      string
	AccessKey   string
	SecretKey   string
	Bucket      string
	QiniuServer string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误！错误提示为：", err)
		os.Exit(1)
	}
	LoadServer(file)
	LoadDB(file)
	LoadQiniu(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":2333")
	JwtKey = file.Section("server").Key("JwtKey").MustString("lindada0911.")
}

func LoadDB(file *ini.File) {
	DBHost = file.Section("database").Key("DBHost").MustString("localhost")
	DBPort = file.Section("database").Key("DBPort").MustString("3306")
	DBUser = file.Section("database").Key("DBUser").MustString("gin")
	DBPassWord = file.Section("database").Key("DBPassWord").MustString("123456")
	DBName = file.Section("database").Key("DBName").MustString("gin")
}

func LoadQiniu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SecretKey = file.Section("qiniu").Key("SecretKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	QiniuServer = file.Section("qiniu").Key("QiniuServer").String()
}
