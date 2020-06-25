package order

import (
	"ny/dorm"
	"ny/stu"
	"strconv"
	"time"
)

func insertOrder(orderType int, stuId int, price string, endTime time.Time, comment string) bool {
	var o order
	o.Comment = comment
	o.FinishTime = &endTime
	o.Type = strconv.Itoa(orderType)
	o.Price = price
	o.StuId = stuId
	o.SchoolId = dorm.GetSchoolId(stu.GetStuDormIdById(o.StuId))
	return o.Insert()
}
