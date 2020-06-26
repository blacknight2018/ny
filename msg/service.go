package msg

import (
	"encoding/json"
	"ny/utils"
)

func InsertTxtMsg(senderStuId int64, recipientStuId int64, content string) bool {
	var m msg
	m.SenderStu = senderStuId
	m.RecipientStu = recipientStuId
	m.Content = content
	m.Type = Txt
	r := m.insert()
	setMsgRead(senderStuId, m.Id)
	return r
}

// 获取A和B之间的最新的limit条聊天记录,并且是A没有阅读过的
func getStuMsg(stuIdA int64, stuIdB int64, limit int) (bool, string) {
	r, d := queryStuMsg(stuIdA, stuIdB, limit)
	if r {

		if bytes, err := json.Marshal(d); err == nil {
			go func() {
				//设置stuIDA对这些消息都阅读了
				for _, v := range d {
					setMsgRead(stuIdA, v.Id)
				}
			}()
			return true, string(bytes)
		}

	}
	return false, utils.EmptyString
}

func setMsgRead(stuId int64, msgId int64) bool {
	var md msgdetail
	md.StuId = stuId
	md.MsgId = msgId
	md.IsRead = true
	return md.insert()
}
