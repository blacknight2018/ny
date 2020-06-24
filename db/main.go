package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetDB() *gorm.DB {
	db, err := gorm.Open("mysql", `root:WOaini@tcp(127.0.0.1:3306)/ny?charset=utf8&parseTime=True&loc=Local&timeout=10s&readTimeout=30s&writeTimeout=60s`)
	db.LogMode(true)
	fmt.Println(err)
	return db
}
