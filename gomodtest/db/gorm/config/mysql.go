package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const (
	USERNAME = "root"
	PASSWORD = "rootroot"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "keepalive"
)

var dsn string = fmt.Sprintf(
	"%s:%s@%s(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
	USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)

func NewMysql() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 开启日志
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
