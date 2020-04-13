package vendors

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Context *gin.Context
}
func (controller Controller) Before() {

}

func (controller Controller) After() {

}

func (controller Controller) Init()  {

}
func (controller Controller) SetContext(context *gin.Context)  {
	controller.Context = context
}

type ControllerInterface interface {
	Init()
	Before()
	After()
	SetContext(*gin.Context)
}

