package main

import (
	"github.com/gin-gonic/gin"
	"ny/dorm"
	"ny/order"
	"ny/school"
	"ny/stu"
	"ny/user"
)

func main() {
	g := gin.Default()
	stu.Test()
	user.Test()
	order.Test()
	user.Register(g)
	dorm.Register(g)
	school.Register(g)
	order.Register(g)
	g.Run(":80")

}
