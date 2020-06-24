package user

import (
	"fmt"
	"ny/db"
)

//user

type user struct {
	Id       int    `gorm:"column:id;unique_index;PRIMARY_KEY"`
	OpenId   string `gorm:"column:open_id;PRIMARY_KEY;"`
	NickName string `gorm:"column:nick_name;"`
	Mobile   string `gorm:"column:mobile;"`
}

func (u *user) TableName() string {
	return "user"
}

func (u *user) insert() bool {
	err := db.GetDB().Model(u).Create(&u).Error
	return err == nil
}

func (u *user) update() bool {
	err := db.GetDB().Model(u).Update(&u).Error
	return err == nil
}

// 根据openId主键信息填充其他字段
func (u *user) queryByOpenId() bool {
	err := db.GetDB().Model(u).Where("open_id = ?", u.OpenId).First(u).Error
	fmt.Println(err)
	return err == nil
}