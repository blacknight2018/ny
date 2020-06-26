package user

import (
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"ny/gerr"
	"ny/stu"
	"ny/utils"
	"strconv"
)

// 展示层函数 只负责接受请求 调用service的接口返回json再返回到前端

func GetUserAvatarUrlByOpenId(OpenId string) (bool, string) {
	return getAvatarUrlByOpenId(OpenId)
}

func GetOpenIdByUserId(userId int64) (bool, string) {
	return getOpenIdById(userId)
}

func SendOrderNotify(openId string, templateId string, dormName string, orderId int64, comment string) bool {
	var param = make(map[string]interface{})
	param["thing2"] = dormName
	param["character_string3"] = strconv.Itoa(int(orderId))
	param["phrase6"] = "已被配送"
	param["thing7"] = comment

	//param["thing2"] = "火龙果等"
	//param["character_string3"] = "AC03704733587501"
	//param["phrase6"] = "下单成功"
	//param["thing7"] = "尊敬的客户,感谢您的支持,请放心"

	return sendNotify(openId, templateId, param)
}
func Register(engine *gin.Engine) {
	g := engine.Group("user")
	g.POST("code", func(context *gin.Context) {
		if ok, data := utils.GetRawData(context); ok {
			code := gjson.Get(data, "code").String()
			if ok, rust := code2Session(code); ok {
				gerr.SetResponse(context, gerr.Ok, &rust)
				return
			}
		}
		gerr.SetResponse(context, gerr.UnKnow, nil)
	})
	g.POST("login", func(context *gin.Context) {
		if ok, data := utils.GetRawData(context); ok {
			openId := gjson.Get(data, "open_id").String()
			nickName := gjson.Get(data, "nick_name").String()
			avatarUrl := gjson.Get(data, "avatar_url").String()
			addUser(openId, nickName, avatarUrl)
			gerr.SetResponse(context, gerr.Ok, nil)
			return
		}
		gerr.SetResponse(context, gerr.UnKnow, nil)
	})
	g.POST("save", func(context *gin.Context) {
		if ok, data := utils.GetRawData(context); ok {
			openId := gjson.Get(data, "open_id").String()
			dormId := gjson.Get(data, "dorm_id").Int()
			stuNumber := gjson.Get(data, "stu_number").String()
			stuMobile := gjson.Get(data, "mobile").String()
			room := gjson.Get(data, "room").String()
			saveMobile(openId, stuMobile)
			if ok, userId := getIdByOpenId(openId); ok {
				ok := stu.SaveDorm(int64(userId), dormId)
				ok2 := stu.SaveStuNumber(int64(userId), stuNumber)
				ok3 := stu.SaveRoom(int64(userId), room)
				if ok && ok2 && ok3 {
					gerr.SetResponse(context, gerr.Ok, nil)
					return
				}

			}

		}
		gerr.SetResponse(context, gerr.UnKnow, nil)
	})
	g.GET("info", func(context *gin.Context) {
		openId := context.Query("open_id")
		if ok, data := getPersonalInfo(openId); ok {
			gerr.SetResponse(context, gerr.Ok, &data)
			return
		}
		gerr.SetResponse(context, gerr.UnKnow, nil)
	})

}

func Test() {
	var u user
	u.OpenId = "aabb"
	//
	//var s stu
	//s.UserId = u.Id
	//s.queryByOpenId()
	//
	//var d dorm
	//d.Id = s.DormId
	//d.queryByOpenId()

	return
}
