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
			"get;/botengine/welcome;VerifyWelcome",
			"post;/botengine/welcome;Welcome",
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

// VerifyWelcome is called when we first verify the webhook for Welcome interaction.
func (c *LiveChatController) VerifyWelcome() {
	fmt.Println("VerifyWelcome!")
	req := c.Ctx.Request()
	err := req.ParseForm()
	if err != nil {
		fmt.Println("[ERR]: " + err.Error())
		c.HTML(http.StatusBadRequest)
		c.Ctx.Template = "error"
		return
	}
	t := req.Form.Get("token")
	if len(t) > 0 && t != "nikkitesttoken" {
		fmt.Println("[ERR]: " + err.Error())
		c.HTML(http.StatusUnauthorized)
		c.Ctx.Template = "error"
		return
	}
	c.Ctx.Write([]byte(req.Form.Get("challenge")))
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
	fmt.Printf("[FORM]: %+v\n", req.Form)
	c.HTML(http.StatusOK)
}
