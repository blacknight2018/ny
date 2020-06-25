package order

import (
	"encoding/json"
	"ny/dorm"
	"ny/stu"
	"ny/utils"
	"strconv"
	"time"
)

func getOrderListBySchoolId(schoolId int) (bool, string) {
	if ok, data := queryListBySchoolId(schoolId); ok {
		if bytes, err := json.Marshal(data); err == nil {
			return true, string(bytes)
		}
	}
	return false, utils.EmptyString
}

func insertOrder(orderType int, stuId int, price string, endTime time.Time, comment string) bool {
	var o order
	o.Comment = comment
	o.FinishTime = &endTime
	o.Type = strconv.Itoa(orderType)
	o.Price = price
	o.StuId = stuId
	o.DormId = stu.GetStuDormIdById(o.StuId)
	o.SchoolId = dorm.GetSchoolId(o.DormId)
	return o.insert()
}
