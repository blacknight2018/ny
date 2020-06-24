package school

import (
	"github.com/gin-gonic/gin"
	"ny/gerr"
)

func Register(engine *gin.Engine) {
	g := engine.Group("school")
	g.GET("list", func(context *gin.Context) {
		if ok, data := querySchoolList(); ok {
			gerr.SetResponse(context, gerr.Ok, &data)
			return
		}
		gerr.SetResponse(context, gerr.UnKnow, nil)
	})
}
