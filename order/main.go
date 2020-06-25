package order

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"ny/gerr"
	"ny/stu"
	"ny/utils"
	"strconv"
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

			if orderType < Delivery || orderType > Buy {
				gerr.SetResponse(context, gerr.ParamError, nil)
				return
			}

			stuId := gjson.Get(data, "stu_id").Int()

			if stu.GetStuExitsById(int(stuId)) == false {
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
	g.GET("/:school_id", func(context *gin.Context) {
		schoolId := context.Param("school_id")
		schoolIdInt, err := strconv.Atoi(schoolId)
		if err == nil {
			if ok, data := getOrderListBySchoolId(schoolIdInt); ok {
				gerr.SetResponse(context, gerr.Ok, &data)
				return
			}
		}
		gerr.SetResponse(context, gerr.UnKnow, nil)
	})
	g.GET("/:school_id/:dorm_id", func(context *gin.Context) {
		schoolId := context.Param("school_id")
		schoolIdInt, err := strconv.Atoi(schoolId)
		dormId := context.Param("dorm_id")
		dormIdInt, err2 := strconv.Atoi(dormId)
		if err == nil && err2 == nil {
			fmt.Println(schoolIdInt, dormIdInt)
			return
		}
		gerr.SetResponse(context, gerr.UnKnow, nil)
	})
	//g.GET("/:school_id/", func(context *gin.Context) {
	//
	//})
}

func Test() {

	//insertOrder(1, "100", time.Now(), "哈哈")
}
