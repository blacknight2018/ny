package user

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"ny/dorm"
	"ny/stu"
	"ny/utils"
	"strings"
)

//服务层 给展示层提供服务返回json的接口

const AppId = "wx53e43951ac45b9a4"
const AppSecret = "3cd2d8c08ebf9f10ec4abeb34e89b828"

type PersonalInfo struct {
	SchoolId  int64  `json:"school_id"`
	DormId    int64  `json:"dorm_id"`
	StuNumber string `json:"stu_number"`
	Mobile    string `json:"mobile"`
	StuId     int64  `json:"stu_id"`
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

func getAccessToken() (bool, string) {
	var param = make(map[string]string)
	param["grant_type"] = "client_credential"
	param["appid"] = AppId
	param["secret"] = AppSecret
	r, data := utils.SendGet("https://api.weixin.qq.com/cgi-bin/token", param)
	if r {
		return true, gjson.Get(data, "access_token").String()
	}
	return false, utils.EmptyString
}

func sendNotify(openId string, templateId string, data map[string]interface{}) bool {
	//UserLogger.InfoLog("sendNotify():" + openId)
	ok, accessToken := getAccessToken()
	if !ok {
		return false
	}
	reqUrl := "https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=" + accessToken

	var param = make(map[string]interface{})
	param["access_token"] = accessToken
	param["touser"] = openId
	param["template_id"] = templateId //"mtm0AlM9wiWsYdE-ihi8lpiioqFi6EtQeusYaY7UrRk"

	//消息参数
	type dataValue struct {
		Value interface{} `json:"value"`
	}
	for k, v := range data {
		data[k] = dataValue{v}
	}
	param["data"] = data

	reqJsonBytes, _ := json.Marshal(param)
	reqJsonString := string(reqJsonBytes)

	req, err1 := http.NewRequest("POST", reqUrl, strings.NewReader(reqJsonString))
	resp, err2 := http.DefaultClient.Do(req)
	if err1 != nil || err2 != nil {
		return false
	}
	defer resp.Body.Close()
	respData, err3 := ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return false
	}
	respDataString := string(respData)
	errCode := int(gjson.Get(respDataString, "errcode").Num)
	if errCode != 0 {
		return false
	}
	return true
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

func getAvatarUrlByOpenId(openId string) (bool, string) {
	var u user
	u.OpenId = openId
	r := u.queryByOpenId()
	return r, u.AvatarUrl
}

func getIdByOpenId(openId string) (bool, int64) {
	var u user
	u.OpenId = openId
	r := u.queryByOpenId()
	return r, u.Id
}

func getOpenIdById(userId int64) (bool, string) {
	var u user
	u.Id = userId
	r := u.queryById()
	return r, u.OpenId
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
