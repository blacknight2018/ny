package user

import (
	"encoding/json"
	"ny/dorm"
	"ny/stu"
	"ny/utils"
)

//服务层 给展示层提供服务返回json的接口

const AppId = "wx53e43951ac45b9a4"
const AppSecret = "3cd2d8c08ebf9f10ec4abeb34e89b828"

type PersonalInfo struct {
	SchoolId  int    `json:"school_id"`
	DormId    int    `json:"dorm_id"`
	StuNumber string `json:"stu_number"`
	Mobile    string `json:"mobile"`
	StuId     int    `json:"stu_id"`
	Room      string `json:"room"`
}

func code2Session(code string) (bool, string) {
	var params = make(map[string]string)
	params["appid"] = AppId
	params["secret"] = AppSecret
	params["js_code"] = code
	params["grant_type"] = code
	return utils.SendGet("https://api.weixin.qq.com/sns/jscode2session", params)
}

func getUserExits(openId string) bool {
	var u user
	u.OpenId = openId
	r := u.queryByOpenId()
	return r
}

func addUser(openId string, nickName string, avatarUrl string) bool {
	var u user
	u.OpenId = openId
	if getUserExits(openId) == false {
		err1 := u.insert()
		if err1 == false {
			return false
		}
	}

	if u.queryByOpenId() && stu.GetStuExitsByUserId(u.Id) == false {
		err, _ := stu.AddStu(u.Id)
		if err == false {
			return false
		}
	}
	u.NickName = nickName
	u.AvatarUrl = avatarUrl
	return u.update()

}
func getIdByOpenId(openId string) (bool, int) {
	var u user
	u.OpenId = openId
	r := u.queryByOpenId()
	return r, u.Id
}

func saveMobile(openId string, mobile string) bool {
	var u user
	u.OpenId = openId
	r := u.queryByOpenId()
	u.Mobile = mobile
	return r && u.update()
}

func getPersonalInfo(openId string) (bool, string) {
	var personInfo PersonalInfo
	var u user
	u.OpenId = openId
	if false == u.queryByOpenId() {
		return false, utils.EmptyString
	}

	var ok, ok1, ok2, ok3, ok4 bool
	ok1, personInfo.DormId = stu.GetStuDormIdByUserId(u.Id)
	ok, personInfo.SchoolId = dorm.GetSchoolId(personInfo.DormId)
	ok2, personInfo.StuNumber = stu.GetStuNumber(u.Id)
	personInfo.Mobile = u.Mobile
	ok3, personInfo.StuId = stu.GetStuIdByUserId(u.Id)
	ok4, personInfo.Room = stu.GetStuRoomByUserId(u.Id)

	bytes, err := json.Marshal(personInfo)
	if err != nil {
		return false, utils.EmptyString
	}
	return ok && ok1 && ok2 && ok3 && ok4, string(bytes)
}
