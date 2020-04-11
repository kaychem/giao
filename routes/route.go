package routes

import (
	"giao/http/controller"
	"giao/vendors"
)

func init()  {
	vendors.App.Route("GET","/", &controller.IndexController{}, "Index")
}
