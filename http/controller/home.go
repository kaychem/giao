package controller

import (
	"fmt"
	"giao/vendors"
)

type IndexController struct {
	vendors.Controller
}
func (controller IndexController) Index()  {
	//context := controller.Context
	//controller.Context.JSON(200, gin.H{
	//	"message": "首页",
	//})
	fmt.Print("首页\n")
}
func (controller IndexController) Hello()  {
	fmt.Print("Hello\n")
}
