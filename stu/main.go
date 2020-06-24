package stu

import (
	"fmt"
	"ny/dorm"
)

func GetStuDormId(userId int) int {
	var s stu
	s.UserId = userId
	s.queryByUserId()
	return s.DormId
}

func GetStuNumber(userId int) string {
	var s stu
	s.UserId = userId
	s.queryByUserId()
	return s.StuNumber
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
	s.queryByUserId()
	s.DormId = dormId
	return s.update()
}

func InsertStu(userId int) (bool, int) {
	var s stu
	s.UserId = userId
	s.DormId = dorm.GetFirstDormId()
	return s.insert(), s.Id
}

func QueryStuExitsByUserId(userId int) bool {
	var s stu
	s.UserId = userId
	r := s.queryByUserId()
	return r
}

func QueryStuIdByUserId(userId int) int {
	var s stu
	s.UserId = userId
	s.queryByUserId()
	return s.Id

}

func QueryStuExitsById(id int) bool {
	var s stu
	s.Id = id
	return s.queryByFirstId()
}

func Test() {
	fmt.Println(QueryStuExitsById(50))
}
