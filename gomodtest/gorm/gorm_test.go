package gorm

import (
	"sync"
	"testing"

	"void.io/gormtry/mysqlsdk"

	"gorm.io/gorm"
)

func TestInitForUT(t *testing.T) {
	type TestTableForUT struct {
		gorm.Model

		X int
		Y string
	}
	dropFunc := mysqlsdk.InitForUT[*TestTableForUT](t, &TestTableForUT{}, []*TestTableForUT{
		{X: 100, Y: "abc"},
		{X: 200, Y: "asdasd"},
		{X: 123, Y: "qweerwt"},
		{X: 566, Y: "zxczc"},
		{X: 820, Y: "jlkjil"},
	}, mysqlsdk.UTOption{
		MysqlDsn: &mysqlsdk.MysqlDsn{
			Username: "root",
			Password: "rootroot",
			Host:     "localhost",
			Port:     "3306",
			DbName:   "test",
		},
		ClearData:  true,
		SoftDelete: false,
		ClearTable: false,
		DropTable:  false,
		InsertOnce: true,
	})
	defer dropFunc()

}

func TestUpdateRowIsLock(t *testing.T) {
	var wg sync.WaitGroup
	var num = 100
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			if err := UpdateRowIsLock(mysqlsdk.DB); err != nil {
				return
			}
		}()
	}

	wg.Wait()
}

func TestOptimisticLock(t *testing.T) {
	type OptimisticLockTest struct {
		gorm.Model

		Name    string
		Num     int64
		Version int64
	}

	obj := &OptimisticLockTest{Name: "phone", Num: 100, Version: 0}

	dropFunc := mysqlsdk.InitForUT[*OptimisticLockTest](t, &OptimisticLockTest{}, []*OptimisticLockTest{obj}, mysqlsdk.UTOption{
		ClearData:  true,
		SoftDelete: false,
		ClearTable: false,
		DropTable:  false,
		InsertOnce: true,
	})
	defer dropFunc()

	mysqlsdk.DB.Transaction(func(tx *gorm.DB) error {
		var val OptimisticLockTest
		if err := tx.Where("id = ?", obj.ID).Find(&val).Error; err != nil {
			t.Fatal(err)
		}
		tx.Where("id = ?", val.ID).
			Where("version = ?", val.Version).
			Updates(&OptimisticLockTest{
				Name:    val.Name,
				Num:     val.Num - 1,
				Version: val.Version + 1,
			})
		return nil
	})
}
