package gorm

import (
	"gorm.io/gorm"
	"log"
)

func UpdateRowIsLock(db *gorm.DB) error {
	exec := db.Exec("update test set a=a+1 where id=1")
	if exec.Error != nil {
		log.Println(exec.Error)
		return exec.Error
	}
	return nil
}
