package msg

import (
	"ny/db"
	"time"
)

type msg struct {
	Id             int64      `gorm:"column:id;PRIMARY_KEY;NOT NULL;"`
	SenderStuId    int64      `gorm:"column:sender_stu;NOT NULL;"`
	RecipientStuId int64      `gorm:"column:recipient_stu;NOT NULL;"`
	Content        string     `gorm:"column:content;NOT NULL;"`
	Type           string     `gorm:"column:type;type:enum('TXT');NOT NULL;"`
	CreateTime     *time.Time `gorm:"column:create_time;"`
}

func (m *msg) TableName() string {
	return "msg"
}

func (m *msg) insert() bool {
	return nil == db.GetDB().Model(m).Create(m).Error
}

// 获取A和B之间的最新的limit条聊天记录,并且是A没有阅读过的
func queryStuMsg(stuIdA int64, stuIdB int64, limit int) (bool, []msg) {
	var m []msg
	sql := `SELECT
	*
FROM
	msg
WHERE
	(
		(sender_stu = ? && recipient_stu = ?) || (sender_stu = ? && recipient_stu = ?)
	) && (
		id NOT IN (
			SELECT
				msg_id
			FROM
				msgdetail
			WHERE
				stu_id = ? && is_read = 1
		)
	) limit ?`
	r := db.GetDB().Model(m).Raw(sql, stuIdA, stuIdB, stuIdB, stuIdA, stuIdA, limit).Find(&m).Error == nil
	return r, m
}
