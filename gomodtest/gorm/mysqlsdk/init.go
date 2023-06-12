package mysqlsdk

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

var internalMysqlDsn MysqlDsn

type MysqlDsn struct {
	Username string
	Password string
	Host     string
	Port     string
	DbName   string
}

func buildDsn(dsn *MysqlDsn) string {
	internalMysqlDsn = *dsn
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		internalMysqlDsn.Username, internalMysqlDsn.Password, internalMysqlDsn.Host, internalMysqlDsn.Port, internalMysqlDsn.DbName)
}

func initGorm(dsn *MysqlDsn, cfg *gorm.Config) {
	cfgVal := &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Info),        // 开启日志
		NamingStrategy: schema.NamingStrategy{SingularTable: true}, // 使用单数表名
	}
	if cfg != nil {
		cfgVal = cfg
	}
	if dsn == nil {
		panic("must set dsn")
	}
	db, err := gorm.Open(mysql.Open(buildDsn(dsn)), cfgVal)
	if err != nil {
		panic(err)
	}
	DB = db
}
