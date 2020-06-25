package dorm

import (
	"github.com/gin-gonic/gin"
	"ny/gerr"
	"strconv"
)

func GetDormName(dormId int) string {
	var d dorm
	d.Id = dormId
	d.QueryById()
	return d.DormName
}

func GetSchoolId(dormId int) (bool, int) {
	var d dorm
	d.Id = dormId

	return d.QueryById(), d.SchoolId
}

func GetFirstDormId() (bool, int) {
	var d dorm
	return d.QueryFirst(), d.Id
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
		if ok, data := getDormList(schoolIdInt); ok {
			gerr.SetResponse(context, gerr.Ok, &data)
			return
		}

	})
}
