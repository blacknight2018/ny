package stu

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ny/dorm"
	"ny/gerr"
	"strconv"
)

func GetStuDormIdByUserId(userId int64) (bool, int64) {
	var s stu
	s.UserId = userId
	return s.queryByUserId(), s.DormId
}

func GetUserIdByStuId(stuId int64) (bool, int64) {
	var s stu
	s.Id = stuId
	r := s.queryById()
	return r, s.UserId
}

func GetStuDormIdById(id int64) (bool, int64) {
	var s stu
	s.Id = id
	return s.queryById(), s.DormId
}

func GetStuNumber(userId int64) (bool, string) {
	var s stu
	s.UserId = userId
	return s.queryByUserId(), s.StuNumber
}

func SaveStuNumber(userId int64, stuNumber string) bool {
	var s stu
	s.UserId = userId
	s.queryByUserId()
	s.StuNumber = stuNumber
	return s.update()
}

func SaveDorm(userId int64, dormId int64) bool {
	var s stu
	s.UserId = userId
	r := s.queryByUserId()
	s.DormId = dormId
	return r && s.update()
}

func SaveRoom(userId int64, room string) bool {
	var s stu
	s.UserId = userId
	r := s.queryByUserId()
	s.DormRoom = room
	return r && s.update()
}

func AddStu(userId int64) (bool, int64) {
	var s stu
	var ok1 bool
	s.UserId = userId
	ok1, s.DormId = dorm.GetFirstDormId()
	return ok1 && s.insert(), s.Id
}

func GetStuExitsByUserId(userId int64) bool {
	var s stu
	s.UserId = userId
	r := s.queryByUserId()
	return r
}

func GetStuIdByUserId(userId int64) (bool, int64) {
	var s stu
	s.UserId = userId

	return s.queryByUserId(), s.Id

}

func GetStuExitsById(id int64) bool {
	var s stu
	s.Id = id
	return s.queryById()
}

func GetStuRoomByUserId(id int64) (bool, string) {
	var s stu
	s.UserId = id
	r := s.queryByUserId()
	return r, s.DormRoom
}

func Register(engine *gin.Engine) {
	g := engine.Group("stu")
	g.GET("", func(context *gin.Context) {
		stuId := context.Query("stu_id")
		if stuIdInt, err := strconv.Atoi(stuId); err == nil {
			if ok, data := getStuDetail(int64(stuIdInt)); ok {
				gerr.SetResponse(context, gerr.Ok, &data)
				return
			}
		}
		gerr.SetResponse(context, gerr.UnKnow, nil)
	})
}

func Test() {
	ok, data := queryStuDetailByUserId(63)
	if ok {
		fmt.Println(data)
	}
}
