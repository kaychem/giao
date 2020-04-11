package controller

import (
	"fmt"
	"giao/vendors"
)

type IndexController struct {
	vendors.Controller
}
func (controller IndexController) Init()  {
	//controller.Context.JSON(200, gin.H{
	//	"message": "消息",
	//})
	fmt.Print("首页")
}
func (controller IndexController) Index()  {
	//controller.Context.JSON(200, gin.H{
	//	"message": "消息",
	//})
	fmt.Print("首页")
}
