package msg

import (
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"ny/gerr"
	"ny/utils"
)

const (
	Txt = "Txt"
)

func Test() {
	//ok, r := getStuMsg(8, 5)
	//setMsgRead(8, 23)
	//ok, r = getStuMsg(8, 5)
	//fmt.Println(ok, r)
}
func Register(engine *gin.Engine) {
	g := engine.Group("msg")
	g.POST("", func(context *gin.Context) {
		if ok, data := utils.GetRawData(context); ok {
			senderStuId := gjson.Get(data, "sender_stuid").Int()
			recipientStuId := gjson.Get(data, "recipient_stuid").Int()
			content := gjson.Get(data, "content").String()
			//fmt.Println(senderStuId, recipientStuId, content)
			if ok := InsertTxtMsg(senderStuId, recipientStuId, content); ok {
				gerr.SetResponse(context, gerr.Ok, nil)
				return
			}
		}
		gerr.SetResponse(context, gerr.UnKnow, nil)
	})
}
