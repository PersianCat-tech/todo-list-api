package conf

import (
	"log"
	"strings"
	"todo_list/model"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	RedisAddr   string
	RedisPw     string
	RedisDbName string
	Db          string
	DbHost      string
	DbPort      string
	DbUser      string
	DbPassWord  string
	DbName      string
)

func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		log.Fatal("cannot load the config file. Please check the file path")
	}
	LoadServer(file)
	LoadMysql(file)
	path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true"}, "")
	model.Database(path) //连接数据库
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()

}

func LoadMysql(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}
