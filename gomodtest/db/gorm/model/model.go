package model

import "gorm.io/gorm"

type TStudent struct {
	gorm.Model
	Id    string `gorm:"primarykey"`
	CarId int
	CId   string
	Name  string
	Sex   string
	Sc    []TStudentCard `gorm:"foreignKey:CarId"`
}

type TStudentCard struct {
	CarId     int `gorm:"primaryKey"`
	Publisher string
}

type Classes struct {
	CId  int
	Name string
}

type User struct {
	gorm.Model
	Username string
	Address  string
	Ord []Order	`gorm:"foreignKey:UserId"`
}

type Order struct {
	gorm.Model
	Number uint
	UserId int
}
