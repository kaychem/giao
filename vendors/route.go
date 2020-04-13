package vendors

func (app *app) Route(method string, relativePath string, controller2 ControllerInterface, action string)  {
	info := &RouteInfo{
		Controller: controller2,
		Path: relativePath,
		Method: method,
		Action :action,
	}
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
