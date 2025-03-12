package model

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Database(conn string) {
	//fmt.Println("conn: ", conn)
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		log.Fatal("database cannot connection")
	}
	db.LogMode(true)

	log.Println("database connect successfully")

	if gin.Mode() == "release" { //在生产模式中使用"release", 减少日志输出，提高性能
		db.LogMode(false)
	}
	db.SingularTable(true)                       //表名不加s， user user
	db.DB().SetMaxIdleConns(20)                  //设置连接池
	db.DB().SetMaxOpenConns(100)                 //设置最大连接数
	db.DB().SetConnMaxLifetime(time.Second * 30) //设置最大连接时间

	DB = db
}
