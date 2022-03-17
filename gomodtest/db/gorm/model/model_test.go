package model

import (
	"fmt"
	"gorm/config"
	"log"
	"testing"
)

func TestStudent(t *testing.T) {
	db, err := config.NewMysql()
	if err != nil {
		log.Fatal(err)
	}
	var stu []TStudent
	//db.Find(&stu)
	//fmt.Println(stu)

	//var s Student
	db.Debug().Preload("TStudentCard").Find(&stu, "180306701")
	fmt.Println(stu)

	//var sc []TStudentCard
	//db.Find(&sc)
	//fmt.Println(sc)
}

func TestUser(t *testing.T) {
	db, err := config.NewMysql()
	if err != nil {
		log.Fatal(err)
	}
	//db.AutoMigrate(&User{}, &Order{})
	var u []User
	db.Preload("Order").Find(&u, 1)
	fmt.Println(u)
}
