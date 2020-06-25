package order

import (
	"ny/db"
	"time"
)

type order struct {
	Id         int        `json:"id";gorm:"column:id;unique_index;PRIMARY_KEY;"`
	StuId      int        `json:"stu_id";gorm:"column:stu_id;"`
	Price      string     `json:"price";gorm:"column:price;"`
	FinishTime *time.Time `json:"finish_time";gorm:"column:finish_time;"`
	Type       string     `json:"type";gorm:"column:type;type:enum('2','1','0')"`
	Comment    string     `json:"comment";gorm:"column:comment;"`
	RecvStu    *int       `json:"recv_stu";gorm:"column:recv_stu;"`
	SchoolId   int        `json:"school_id";gorm:"column:school_id;"`
}

func (o *order) TableName() string {
	return "order"
}

// 更新非空字段内容
func (o *order) Update() bool {
	err := db.GetDB().Model(o).Update(o).Error
	return err == nil
}

func (o *order) Insert() bool {
	err := db.GetDB().Model(o).Create(&o).Error
	return err == nil
}
