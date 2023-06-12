package mysqlsdk

import (
	"testing"

	"gorm.io/gorm"
)

type UTOption struct {
	GormCfg  *gorm.Config
	MysqlDsn *MysqlDsn

	// 优先级：DropTable > ClearTable > ClearData
	ClearData  bool // 测试结束后，是否删除本次测试的数据
	ClearTable bool // 测试结束后，是否清空表内所有数据
	DropTable  bool // 测试结束后，是否删除表
	InsertOnce bool // true => 如果表内已经有数据，则不插入测试数据
	SoftDelete bool // true => 软删除
}

func InitForUT[T any](t *testing.T, table T, testData []T, opt UTOption) (drop func() error) {
	initGorm(opt.MysqlDsn, opt.GormCfg)
	if err := DB.AutoMigrate(table); err != nil {
		t.Fatal(err)
	}

	if opt.InsertOnce {
		var count int64
		if err := DB.Model(table).Count(&count).Error; err != nil {
			t.Fatal(err)
		}
		// 如果表中不存在数据，则插入数据
		if count == 0 {
			if err := DB.Create(&testData).Error; err != nil {
				t.Fatal(err)
			}
		}
	} else {
		if err := DB.Create(&testData).Error; err != nil {
			t.Fatal(err)
		}
	}

	return func() error {
		switch {
		case opt.DropTable:
			return DB.Migrator().DropTable(table)
		case opt.ClearTable:
			if !opt.SoftDelete {
				return DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Unscoped().Delete(table).Error
			}
			return DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(table).Error
		case opt.ClearData:
			if !opt.SoftDelete {
				return DB.Unscoped().Delete(testData).Error
			}
			return DB.Delete(testData).Error
		}
		return nil
	}
}
