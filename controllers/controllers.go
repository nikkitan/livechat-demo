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
			"get;/botengine/prohibiteditems;ProhibitedItems",
			"get;/botengine/welcome;Welcome",
		},
	}
}

//Home is responsible for rendering home page.
func (c *LiveChatController) Home() {
	fmt.Println("Home!")
	c.Ctx.Template = "index"
	c.HTML(http.StatusOK)
}

// ProhibitedItems is webhook for when the user asks about prohibited items.
func (c *LiveChatController) ProhibitedItems() {
	fmt.Println("Welcome!")
	c.Ctx.Template = "index"
	c.HTML(http.StatusOK)
}

// Welcome is responsible for getting user info when the Bot first welcomed the user.
func (c *LiveChatController) Welcome() {
	fmt.Println("Welcome!")
	req := c.Ctx.Request()
	err := req.ParseForm()
	if err != nil {
		fmt.Println("[ERR]: " + err.Error())
		c.HTML(http.StatusBadRequest)
		c.Ctx.Template = "error"
		return
	}
	t := req.Form.Get("token")
	if t != "nikkitesttoken" {
		fmt.Println("[ERR]: " + err.Error())
		c.HTML(http.StatusUnauthorized)
		c.Ctx.Template = "error"
		return
	}
	c.Ctx.Write([]byte("challenge"))
	c.HTML(http.StatusOK)
}
