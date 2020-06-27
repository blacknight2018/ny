package stu

import (
	"encoding/json"
	"ny/utils"
)

type Detail struct {
	StuId     int64  `json:"stu_id"`
	AvatarUrl string `json:"avatar_url"`
}

// 获取学生的信息，用于展示
func getStuDetail(StuId int64) (bool, string) {
	ok, data := queryStuDetailByStuId(StuId)
	if ok {
		if bytes, err := json.Marshal(data); err == nil {
			return true, string(bytes)
		}
	}
	return false, utils.EmptyString
}
