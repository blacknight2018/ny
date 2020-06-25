package stu

import "ny/db"

//stu

type stu struct {
	Id        int    `gorm:"column:id;unique_index;PRIMARY_KEY;"`
	DormId    int    `gorm:"column:dorm_id;"`
	StuNumber string `gorm:"column:stu_number;"`
	UserId    int    `gorm:"column:user_id;"`
	DormRoom  string `gorm:"column:dorm_room;"`
}

func (s *stu) TableName() string {
	return "stu"
}

// 更新非空字段内容
func (s *stu) update() bool {
	err := db.GetDB().Model(s).Update(&s).Error
	return err == nil
}

// 根据主键信息填充其他字段
func (s *stu) queryByUserId() bool {
	err := db.GetDB().Model(s).Where("user_id = ?", s.UserId).First(s).Error
	return err == nil
}

//

func (s *stu) queryById() bool {
	err := db.GetDB().Model(s).Where("id = ?", s.Id).First(s).Error
	return err == nil
}

// 插入一行
func (s *stu) insert() bool {
	err := db.GetDB().Model(s).Create(s).Error
	return err == nil
}
