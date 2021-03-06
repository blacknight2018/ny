package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var gdb *gorm.DB

func GetDB() *gorm.DB {
	if gdb != nil {
		return gdb
	}
	db, err := gorm.Open("mysql", `root:WOaini@tcp(127.0.0.1:3306)/ny?charset=utf8&parseTime=True&loc=Local&timeout=10s&readTimeout=30s&writeTimeout=60s`)
	db.LogMode(true)
	gdb = db
	fmt.Println(err)
	return db
}
