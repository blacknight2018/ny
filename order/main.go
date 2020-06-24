package order

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"ny/gerr"
	"ny/stu"
	"ny/utils"
	"time"
)

const (
	Delivery = iota
	Food     = iota
	Buy      = iota
)

func Register(engine *gin.Engine) {
	g := engine.Group("order")
	g.POST("", func(context *gin.Context) {
		if ok, data := utils.GetRawData(context); ok {
			orderType := gjson.Get(data, "type").Int()

			if orderType < 0 || orderType > Buy {
				gerr.SetResponse(context, gerr.ParamError, nil)
				return
			}

			stuId := gjson.Get(data, "stu_id").Int()

			if stu.QueryStuExitsById(int(stuId)) == false {
				gerr.SetResponse(context, gerr.UnKnowUser, nil)
				return
			}

			orderPrice := gjson.Get(data, "price").String()
			orderEndTime := gjson.Get(data, "end_time").Int()
			orderComment := gjson.Get(data, "comment").String()
			fmt.Println(orderType, orderPrice, orderEndTime, orderComment)

			if insertOrder(int(orderType), int(stuId), orderPrice, time.Unix(orderEndTime/1000, 0), orderComment) {
				gerr.SetResponse(context, gerr.Ok, nil)
				return
			}

		}
		gerr.SetResponse(context, gerr.UnKnow, nil)
	})
}

func Test() {

	//insertOrder(1, "100", time.Now(), "哈哈")
}
