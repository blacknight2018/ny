package msg

import "ny/db"

type msgdetail struct {
	Id     int64 `gorm:"column:id;PRIMARY_KEY;NOT NULL;"`
	MsgId  int64 `gorm:"column:msg_id;NUL NULL;"`
	StuId  int64 `gorm:"column:stu_id;NOT NULL;"`
	IsRead bool  `json:"is_read,omitempty";gorm:"column:is_read;NOT NULL;"`
}

func (md *msgdetail) TableName() string {
	return "msgdetail"
}

func (md *msgdetail) insert() bool {
	return nil == db.GetDB().Model(md).Create(md).Error
}
