package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var(
	db * gorm.DB
)

func connect(){
	d,err := gorm.Open("postgres", "postgres:csigesnpb")

	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}