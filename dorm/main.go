package dorm

import (
	"github.com/gin-gonic/gin"
	"ny/gerr"
	"strconv"
)

func GetDormName(dormId int64) (bool, string) {
	var d dorm
	d.Id = dormId
	r := d.queryById()
	return r, d.DormName
}

func GetSchoolId(dormId int64) (bool, int64) {
	var d dorm
	d.Id = dormId

	return d.queryById(), d.SchoolId
}

func GetFirstDormId() (bool, int64) {
	var d dorm
	return d.queryFirst(), d.Id
}

func Register(engine *gin.Engine) {
	g := engine.Group("dorm")
	g.GET("/:school_id", func(context *gin.Context) {
		schoolId := context.Param("school_id")
		schoolIdInt, err := strconv.Atoi(schoolId)
		if err != nil {
			gerr.SetResponse(context, gerr.UnKnow, nil)
			return
		}
		if ok, data := getDormList(int64(schoolIdInt)); ok {
			gerr.SetResponse(context, gerr.Ok, &data)
			return
		}

	})
}
