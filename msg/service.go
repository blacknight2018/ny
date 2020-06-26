package msg

func InsertTxtMsg(senderStuId int64, recipientStuId int64, content string) bool {
	var m msg
	m.SenderStuId = senderStuId
	m.RecipientStuId = recipientStuId
	m.Content = content
	m.Type = Txt
	return m.insert()
}

func getStuMsg(stuId int64, limit int) (bool, []msg) {
	return queryStuMsg(stuId, limit)
}

func setMsgRead(stuId int64, msgId int64) bool {
	var md msgdetail
	md.StuId = stuId
	md.MsgId = msgId
	md.IsRead = true
	return md.insert()
}
