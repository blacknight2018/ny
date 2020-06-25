package user

import (
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"ny/gerr"
	"ny/stu"
	"ny/utils"
)

// 展示层函数 只负责接受请求 调用service的接口返回json再返回到前端

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
				ok := stu.SaveDorm(userId, int(dormId))
				ok2 := stu.SaveStuNumber(userId, stuNumber)
				ok3 := stu.SaveRoom(userId, room)
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
