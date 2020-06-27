package stu

import "ny/db"

type stuDetail struct {
	UserId    int64  `json:"user_id";gorm:"column:user_id;"`
	AvatarUrl string `json:"avatar_url";gorm:"column:avatar_url;"`
	NickName  string `json:"nick_name";gorm:"column:nick_name;"`
	StuId     int64  `json:"stu_id";gorm:"column:stu_id;"`
}

func queryStuDetailByUserId(userId int64) (bool, stuDetail) {
	var r stuDetail
	sql := `SELECT
	USER .id as user_id,
	USER .avatar_url,
	USER .nick_name,
	stu.id as stu_id
FROM
	USER,
	stu
WHERE
	USER .id = ? && stu.user_id = ?`
	return nil == db.GetDB().Raw(sql, userId, userId).First(&r).Error, r
}

func queryStuDetailByStuId(stuId int64) (bool, stuDetail) {
	var r stuDetail
	sql := `SELECT
	USER .id as user_id,
	USER .avatar_url,
	USER .nick_name,
	stu.id as stu_id
FROM
	USER,
	stu
WHERE
	stu.id = ?`
	return nil == db.GetDB().Raw(sql, stuId).First(&r).Error, r
}
