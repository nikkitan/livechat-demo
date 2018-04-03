package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

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
			"get;/livechat/livechat_home;Home",
			"get;/livechat/prohibiteditems;ProhibitedItems",
			"get;/livechat/welcome;VerifyWebhook",
			"post;/livechat/welcome;Welcome",
			"get;/livechat/appcrashing;VerifyWebhook",
			"post;/livechat/appcrashing;AppCrashing",
		},
	}
}

//Home is responsible for rendering home page.
func (c *LiveChatController) Home() {
	fmt.Println("Home!")
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

	c.Ctx.Template = "livechat_home"
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

	type p struct {
		ItemID string `json:"itemid"`
		Name   string `json:"name"`
		Email  string `json:"email"`
	}

	type response struct {
		Type     string   `json:"type"`
		Elements []string `json:"elements"`
	}

	var result struct {
		Responses  []response `json:"responses"`
		Parameters p          `json:"parameters"`
	}

	w.Header().Set("Content-Type", "application/json")
	result.Responses = []response{
		{
			Type:     "text",
			Elements: []string{"Is your name {{name}}?"},
		},
	}
	result.Parameters.ItemID = "1234567890"
	result.Parameters.Name = "nikkitan"
	result.Parameters.Email = "nikkitan+222@mercari.com"

	fmt.Printf("[JSON]: %+v.\n", result)

	json.NewEncoder(os.Stdout).Encode(result)

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		fmt.Printf("[JSON_ERR]: %s\n", err.Error())
	}

	fmt.Println("[DONE_JSON]")

	fmt.Printf("[RESP]: %+v\n", c.Ctx.Response())

}

// ProhibitedItems is webhook for when the user asks about prohibited items.
func (c *LiveChatController) ProhibitedItems() {
	fmt.Println("Welcome!")
	c.Ctx.Template = "livechat_home"
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
