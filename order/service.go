package order

import (
	"encoding/json"
	"fmt"
	"ny/dorm"
	"ny/stu"
	"ny/user"
	"ny/utils"
	"strconv"
	"time"
)

func getOrderListBySchoolId(schoolId int64) (bool, string) {
	if ok, data := queryListBySchoolId(schoolId); ok {
		if bytes, err := json.Marshal(data); err == nil {
			return true, string(bytes)
		}
	}
	return false, utils.EmptyString
}
func getOrderListBySchoolIdDormId(schoolId int64, dormId int64) (bool, string) {
	if ok, data := queryListBySchoolIdDormId(schoolId, dormId); ok {
		if bytes, err := json.Marshal(data); err == nil {
			return true, string(bytes)
		}
	}
	return false, utils.EmptyString
}

func finishOrder(openId string, orderId int64, stuId int64) bool {
	var o order
	o.Id = orderId
	r := o.queryById()

	ok, dormName := dorm.GetDormName(o.DormId)
	if !r || !ok {
		return false
	}
	fmt.Println(dormName)
	//已有人完成，不可再设置完成
	if o.RecvStu != nil {
		return true
	}

	//设置完成的stu_id
	o.RecvStu = &stuId
	if !o.update() {
		return false
	}

	//发送通知
	return user.SendOrderNotify(openId, o.TemplateId, dormName, orderId, o.Comment)

}

func insertOrder(orderType int, stuId int64, price string, endTime time.Time, comment string, templateId string) bool {
	var o order
	var ok, ok2, ok3, ok4, ok5 bool
	var userId int64
	var openId string
	o.Comment = comment
	o.FinishTime = &endTime
	o.Type = strconv.Itoa(orderType)
	o.Price = price
	o.StuId = stuId
	o.TemplateId = templateId
	ok, o.DormId = stu.GetStuDormIdById(o.StuId)
	ok2, o.SchoolId = dorm.GetSchoolId(o.DormId)
	ok3, userId = stu.GetUserIdByStuId(o.StuId)
	ok4, openId = user.GetOpenIdByUserId(userId)
	ok5, o.AvatarUrl = user.GetUserAvatarUrlByOpenId(openId)

	return ok && ok2 && ok3 && ok4 && ok5 && o.insert()
}
