package controllers

import (
	"bytes"
	"fmt"
	"io/ioutil"
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
			"get;/botengine/welcome;VerifyWebhook",
			"post;/botengine/welcome;Welcome",
			"get;/botengine/appcrashing;VerifyWebhook",
			"post;/botengine/appcrashing;AppCrashing",
		},
	}
}

//Home is responsible for rendering home page.
func (c *LiveChatController) Home() {
	fmt.Println("Home!")
	c.Ctx.Template = "index"
	c.HTML(http.StatusOK)
}

// AppCrashing is webhook for when the user asks about prohibited items.
func (c *LiveChatController) AppCrashing() {
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
	fmt.Printf("[CTX]: %+v\n", c.Ctx)
	fmt.Printf("[RESP]: %+v\n", c.Ctx.Data)

}

// ProhibitedItems is webhook for when the user asks about prohibited items.
func (c *LiveChatController) ProhibitedItems() {
	fmt.Println("Welcome!")
	c.Ctx.Template = "index"
	c.HTML(http.StatusOK)
}

// VerifyWebhook is called when we first verify a webhook for registering it.
func (c *LiveChatController) VerifyWebhook() {
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
	fmt.Printf("[HEADER]: %+v\n", req.Header)
	fmt.Printf("[FORM]: %+v\n", req.Form)
	c.HTML(http.StatusOK)
}
