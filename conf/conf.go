package conf

import (
	"log"

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
	file, err := ini.Load("./conf.ini")
	if err != nil {
		log.Fatal("cannot load the config file. Please check the file path")
	}
	
}
