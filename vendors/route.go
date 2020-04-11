package vendors

func (app *app) Route(method string, relativePath string, controller2 ControllerInterface, action string)  {
	info := &RouteInfo{}
	info.Controller = controller2
	info.Path = relativePath
	info.Method = method
	info.Action = action
	app.Routes = append(app.Routes, info)
}

type RouteInfo struct {
	Method      string
	Path        string
	Handler     string
	Action      string
	Controller ControllerInterface
}

type HandlerFunc func()
