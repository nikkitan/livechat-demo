package controllers

import (
	"fmt"
	"net/http"

	"github.com/gernest/utron/controller"
)

// LiveChatController is for controllers of livechat-demo.
type LiveChatController struct {
	controller.BaseController
	Routes []string
}

// NewLiveChatController is the constructor of LivChatController.
func NewLiveChatController() controller.Controller {
	return &LiveChatController{
		Routes: []string{
			"get;/;Home",
		},
	}
}

//Home is responsible for rendering home page.
func (c *LiveChatController) Home() {
	fmt.Println("Home!")
	c.HTML(http.StatusOK)
}
