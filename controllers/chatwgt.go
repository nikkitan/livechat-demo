package controllers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gernest/utron/controller"
)

// ChatWgtController is for controllers of livechat-demo.
type ChatWgtController struct {
	controller.BaseController
	Routes []string
}

// NewChatWgtController is the constructor of LivChatController.
func NewChatWgtController() controller.Controller {
	return &ChatWgtController{
		Routes: []string{
			"get;/chatwgt/chatwgt_home;Home",
			"get;/chatwgt/prohibiteditems;ProhibitedItems",
			"get;/chatwgt/welcome;VerifyWebhook",
			"post;/chatwgt/welcome;Welcome",
			"get;/chatwgt/appcrashing;VerifyWebhook",
			"post;/chatwgt/appcrashing;AppCrashing",
		},
	}
}

//Home is responsible for rendering home page.
func (c *ChatWgtController) Home() {
	fmt.Println("Home!")
	c.Ctx.Template = "chatwgt_home"
	c.HTML(http.StatusOK)
}

// AppCrashing is webhook for when the user asks about prohibited items.
func (c *ChatWgtController) AppCrashing() {
	fmt.Println("AppCrashing!")
	req := c.Ctx.Request()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("Error reading body: %v", err)
		return
	}

	// Work / inspect body. You may even modify it!

	// And now set a new body, which will simulate the same data we read:
	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	err = req.ParseForm()
	if err != nil {
		fmt.Println("[ERR]: " + err.Error())
		c.HTML(http.StatusBadRequest)
		c.Ctx.Template = "error"
		return
	}
	fmt.Printf("[HEADER]: %+v\n", req.Header)
	fmt.Printf("[FORM]: %+v\n", req.Form)
	fmt.Printf("[REQ]: %+v\n", req)
	fmt.Printf("[REQ_BODY]: %+v\n", req.Body)

	c.HTML(http.StatusOK)
	w := c.Ctx.Response()

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"test\":\"test\"}"))
	fmt.Printf("[RESP]: %+v\n", c.Ctx.Response())

}

// ProhibitedItems is webhook for when the user asks about prohibited items.
func (c *ChatWgtController) ProhibitedItems() {
	fmt.Println("Welcome!")
	c.Ctx.Template = "index"
	c.HTML(http.StatusOK)
}

// VerifyWebhook is called when we first verify a webhook for registering it.
func (c *ChatWgtController) VerifyWebhook() {
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
func (c *ChatWgtController) Welcome() {
	fmt.Println("Welcome!")
	req := c.Ctx.Request()
	err := req.ParseForm()
	if err != nil {
		fmt.Println("[ERR]: " + err.Error())
		c.HTML(http.StatusBadRequest)
		c.Ctx.Template = "error"
		return
	}
	fmt.Printf("[HEADER]: %+v\n", req.Header)
	fmt.Printf("[FORM]: %+v\n", req.Form)
	c.HTML(http.StatusOK)
}