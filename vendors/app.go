package vendors

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

type app struct {
	Engine *gin.Engine
	Routes []*RouteInfo
}
var (
	// app 是一个实例
	App *app
)

func init()  {
	App = New()
}
func (app *app) Start(addr string)  {
	app.handleRoute()
	app.run(addr)
}
func New() *app {
	engine := &app{Engine: gin.Default()}
	return engine
}
func (app *app) run(addr string)  {
	App.Engine.Run(addr) // 监听并在 0.0.0.0:8080 上启动服务
}

func (app *app) handleRoute()  {
	for _, route := range app.Routes {
		method := route.Method
		action := route.Action

		vc := reflect.New(
				reflect.Indirect(reflect.ValueOf(route.Controller)).Type())

		in := make([]reflect.Value, 0)

		App.Engine.Handle(method, route.Path, func(context *gin.Context) {
			//vc.FieldByName("Context") =  context
			route.Controller.Init()
			route.Controller.Before()
			vc.MethodByName(action).Call(in)
			route.Controller.After()
		})
	}
}
