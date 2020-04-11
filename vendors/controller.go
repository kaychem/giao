package vendors

import "github.com/gin-gonic/gin"

type Controller struct {
	Context *gin.Context
}

func (controller Controller) Init()  {

}

type ControllerInterface interface {
	Init()
	Before()
	After()
}

