package main

import (
	"github.com/gin-gonic/gin"
	"ny/dorm"
	"ny/msg"
	"ny/order"
	"ny/school"
	"ny/stu"
	"ny/user"
)

func main() {
	g := gin.Default()
	msg.Test()
	stu.Test()
	user.Test()
	order.Test()
	user.Register(g)
	dorm.Register(g)
	school.Register(g)
	order.Register(g)
	msg.Register(g)
	g.Run(":80")

}
