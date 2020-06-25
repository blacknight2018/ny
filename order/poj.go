package order

import (
	"ny/db"
	"time"
)

type order struct {
	Id         int        `json:"id";gorm:"column:id;unique_index;PRIMARY_KEY;"`
	StuId      int        `json:"stu_id";gorm:"column:stu_id;NOT NULL;"`
	Price      string     `json:"price";gorm:"column:price;"`
	FinishTime *time.Time `json:"finish_time";gorm:"column:finish_time;"`
	Type       string     `json:"type";gorm:"column:type;type:enum('2','1','0');NOT NULL"`
	Comment    string     `json:"comment";gorm:"column:comment;"`
	RecvStu    *int       `json:"recv_stu,omitempty";gorm:"column:recv_stu;"`
	SchoolId   int        `json:"school_id";gorm:"column:school_id;NOT NULL;"`
	DormId     int        `json:"dorm_id";gorm:"column:dorm_id;NOT NULL;"`
}

func (o *order) TableName() string {
	return "order"
}

// 更新非空字段内容
func (o *order) update() bool {
	err := db.GetDB().Model(o).Update(o).Error
	return err == nil
}

func (o *order) insert() bool {
	err := db.GetDB().Model(o).Create(&o).Error
	return err == nil
}

func queryListBySchoolId(schoolId int) (bool, []order) {
	var os []order
	err := db.GetDB().Model(os).Where("school_id = ?", schoolId).Find(&os).Error
	return err == nil, os

}
