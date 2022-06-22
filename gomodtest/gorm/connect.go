package gorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var mysqlConn *gorm.DB

const (
	USERNAME = "root"
	PASSWORD = "rootroot"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "test"
)

var dsn = fmt.Sprintf(
	"%s:%s@%s(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
	USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)

func init() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 开启日志
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		panic(err)
	}
	mysqlConn = db
}
