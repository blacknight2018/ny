package stu

import (
	"fmt"
	"ny/dorm"
)

func GetStuDormIdByUserId(userId int) (bool, int) {
	var s stu
	s.UserId = userId
	return s.queryByUserId(), s.DormId
}

func GetUserIdByStuId(stuId int) (bool, int) {
	var s stu
	s.Id = stuId
	r := s.queryById()
	return r, s.UserId
}

func GetStuDormIdById(id int) (bool, int) {
	var s stu
	s.Id = id
	return s.queryById(), s.DormId
}

func GetStuNumber(userId int) (bool, string) {
	var s stu
	s.UserId = userId
	return s.queryByUserId(), s.StuNumber
}

func SaveStuNumber(userId int, stuNumber string) bool {
	var s stu
	s.UserId = userId
	s.queryByUserId()
	s.StuNumber = stuNumber
	return s.update()
}

func SaveDorm(userId int, dormId int) bool {
	var s stu
	s.UserId = userId
	r := s.queryByUserId()
	s.DormId = dormId
	return r && s.update()
}

func SaveRoom(userId int, room string) bool {
	var s stu
	s.UserId = userId
	r := s.queryByUserId()
	s.DormRoom = room
	return r && s.update()
}

func AddStu(userId int) (bool, int) {
	var s stu
	var ok1 bool
	s.UserId = userId
	ok1, s.DormId = dorm.GetFirstDormId()
	return ok1 && s.insert(), s.Id
}

func GetStuExitsByUserId(userId int) bool {
	var s stu
	s.UserId = userId
	r := s.queryByUserId()
	return r
}

func GetStuIdByUserId(userId int) (bool, int) {
	var s stu
	s.UserId = userId

	return s.queryByUserId(), s.Id

}

func GetStuExitsById(id int) bool {
	var s stu
	s.Id = id
	return s.queryById()
}

func GetStuRoomByUserId(id int) (bool, string) {
	var s stu
	s.UserId = id
	r := s.queryByUserId()
	return r, s.DormRoom
}

func Test() {
	fmt.Println(GetStuExitsById(50))
}
