package controllers

import (
	"net/http"

	"github.com/gernest/utron/controller"
)

// CommonController is for controllers that can be used for all different GUI-system specific controllers.
type CommonController struct {
	controller.BaseController
	Routes []string
}

// NewCommonController is the constructor of CommonController.
func NewCommonController() controller.Controller {
	return &CommonController{
		Routes: []string{
			"get;/;Home",
		},
	}
}

//Home is responsible for rendering home page.
func (c *CommonController) Home() {
	c.Ctx.Template = "index"
	c.HTML(http.StatusOK)
}
