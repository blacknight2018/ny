package msg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"ny/gerr"
	"ny/utils"
	"strconv"
)

const (
	Txt = "Txt"
)

func Test() {
	ok, r := getStuMsg(9, 8, 5)
	setMsgRead(9, 22)
	ok, r = getStuMsg(9, 8, 5)
	fmt.Println(ok, r)

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
	g.GET("", func(context *gin.Context) {
		stuIdA := context.Query("stua_id")
		stuIdB := context.Query("stub_id")
		stuIdAInt, err := strconv.Atoi(stuIdA)
		stuIdBInt, err2 := strconv.Atoi(stuIdB)
		if err == nil && err2 == nil {
			fmt.Println(stuIdAInt, stuIdBInt)
			if ok, data := getStuMsg(int64(stuIdAInt), int64(stuIdBInt), 5); ok {
				gerr.SetResponse(context, gerr.Ok, &data)
			}
			return
		}
		gerr.SetResponse(context, gerr.UnKnow, nil)

	})
}
