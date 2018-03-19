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
			"get;/chatwgt/getpurchaseditems;VerifyWebhook",
			"post;/chatwgt/getpurchaseditems;GetPurchasedItemsAsCards",
			"get;/chatwgt/getsolditems;VerifyWebhook",
			"post;/chatwgt/getsolditems;GetSoldItemsAsCards",
			"get;/chatwgt/other;VerifyWebhook",
			"post;/chatwgt/other;Other",
			"get;/chatwgt/itemops;VerifyWebhook",
			"post;/chatwgt/itemops;GenerateItemOpQuickReplies",
			"get;/chatwgt/itemstat;VerifyWebhook",
			"post;/chatwgt/itemstat;GetItemStatus",
			"get;/chatwgt/return;VerifyWebhook",
			"post;/chatwgt/return;RequestReturn",
			"get;/chatwgt/cancel;VerifyWebhook",
			"post;/chatwgt/cancel;CancelItem",
		},
	}
}

//Home is responsible for rendering home page.
func (c *ChatWgtController) Home() {
	c.Ctx.Template = "chatwgt_home"
	c.HTML(http.StatusOK)
}

//GetItemStatus is responsible for rendering home page.
func (c *ChatWgtController) GetItemStatus() {
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

	w := c.Ctx.Response()

	//TODO: SYNCHRONOUS call to Mercari backend for the items of current user.
	type response struct {
		Type     string   `json:"type"`
		Elements []string `json:"elements"`
	}

	var result struct {
		Responses []response `json:"responses"`
	}

	w.Header().Set("Content-Type", "application/json")
	result.Responses = []response{
		{
			Type: "text",
			Elements: []string{
				"Your order is on its way.\n Scheduled to arrive on 02/01/2018.",
			},
		}, {
			Type: "text",
			Elements: []string{
				"Here is the tracking information: http://fedex.com/....",
			},
		},
	}
	fmt.Printf("[JSON]: %+v.\n", result)

	json.NewEncoder(os.Stdout).Encode(result)

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		fmt.Printf("[JSON_ERR]: %s\n", err.Error())
	}

	fmt.Println("[DONE_JSON]")

	fmt.Printf("[RESP]: %+v\n", c.Ctx.Response())

	c.HTML(http.StatusOK)
}

// AppCrashing is webhook for when the user asks about prohibited items.
func (c *ChatWgtController) AppCrashing() {
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

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"test\":\"test\"}"))
	fmt.Printf("[RESP]: %+v\n", c.Ctx.Response())

}

// VerifyWebhook is called when we first verify a webhook for registering it.
func (c *ChatWgtController) VerifyWebhook() {
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

// GenerateItemOpQuickReplies constructs and returns a BotEngine's Quick-Replies GUI obj
// with proper data and wiring with other iteractions.
// Request Parameters:  item ID.
// Return: A Quick-Replies object for the bot to display.
func (c *ChatWgtController) GenerateItemOpQuickReplies() {
	fmt.Println("GenerateItemOpQuickReplies!")
	req := c.Ctx.Request()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("Error reading body: %v", err)
		return
	}
	//TODO: Parse the JSON in body.

	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	err = req.ParseForm()

	if err != nil {
		fmt.Println("[ERR]: " + err.Error())
		c.HTML(http.StatusBadRequest)
		c.Ctx.Template = "error"
		return
	}

	//Generate the Quick Replies.
	type p struct {
		ItemID string `json:"current_item_id"`
	}

	type response struct {
		Type    string   `json:"type"`
		Title   string   `json:"title"`
		Buttons []Button `json:"buttons"`
	}

	var result struct {
		Responses  []response `json:"responses"`
		Parameters p          `json:"parameters"`
	}
	w := c.Ctx.Response()

	w.Header().Set("Content-Type", "application/json")

	result.Responses = []response{
		{
			Type:  "quickReplies",
			Title: "Please pick what you want to do for the item:",
			Buttons: []Button{
				{
					Type:  Goto,
					Title: "Item Status",
					Value: "5aa833dcf60bd80007b25375",
				}, {
					Type:  Goto,
					Title: "Cancel Item",
					Value: "5aa83bcbf60bd80007b253e3",
				}, {
					Type:  Goto,
					Title: "Return Item",
					Value: "5aa83bbd7eefe000078cb066",
				},
			},
		},
	}

	result.Parameters.ItemID = "fakeitem1111"

	fmt.Printf("[JSON]: %+v.\n", result)

	json.NewEncoder(os.Stdout).Encode(result)

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		fmt.Printf("[JSON_ERR]: %s\n", err.Error())
	}

	fmt.Println("[DONE_JSON]")
	fmt.Printf("[RESP]: %+v\n", c.Ctx.Response())
	c.HTML(http.StatusOK)

}

// GetSoldItemsAsCards gets the purchased or sold items for a user.
// Request Parameters:  user's first name
//						user's last name
//						user's email.
// Return: A Cards object for the bot to display.
func (c *ChatWgtController) GetSoldItemsAsCards() {
	fmt.Println("GetSoldItemsAsCards!")
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
	fmt.Printf("[REQ_BODY]: %+v\n", req.Body)

	//TODO: SYNCHRONOUS call to Mercari backend for the items of current user.

	// Return the list of items as Cards.
	type p struct {
		ItemID string `json:"itemid"`
		Name   string `json:"name"`
		Email  string `json:"email"`
	}

	type response struct {
		Type     string `json:"type"`
		Elements []Card `json:"elements"`
	}

	var result struct {
		Responses []response `json:"responses"`
	}

	w := c.Ctx.Response()

	w.Header().Set("Content-Type", "application/json")

	result.Responses = []response{
		{
			Type: "cards",
			Elements: []Card{
				{
					Title:    "Mario Chess",
					ImageURL: "https://image.ibb.co/hYWMXx/mariochess.jpg",
					Buttons: []Button{
						{
							Type:  Postback,
							Title: "Item Operations",
							Value: "fakeitemid1111",
						},
					},
				}, {
					Title:    "Nike Air",
					ImageURL: "https://image.ibb.co/iUGoCx/nike.jpg",
					Buttons: []Button{
						{
							Type:  Postback,
							Title: "Item Operations",
							Value: "fakeitemid2222",
						},
					},
				},
			},
		},
	}

	fmt.Printf("[JSON]: %+v.\n", result)

	json.NewEncoder(os.Stdout).Encode(result)

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		fmt.Printf("[JSON_ERR]: %s\n", err.Error())
	}

	fmt.Println("[DONE_JSON]")
	fmt.Printf("[RESP]: %+v\n", c.Ctx.Response())
	c.HTML(http.StatusOK)
}

// GetPurchasedItemsAsCards gets the purchased or sold items for a user.
// Request Parameters:  user's first name
//						user's last name
//						user's email.
// Return: A Cards object for the bot to display.
func (c *ChatWgtController) GetPurchasedItemsAsCards() {
	fmt.Println("GetPurchasedItemsAsCards!")
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
	fmt.Printf("[REQ_BODY]: %+v\n", req.Body)

	//TODO: SYNCHRONOUS call to Mercari backend for the items of current user.

	// Return the list of items as Cards.
	type p struct {
		ItemID string `json:"itemid"`
		Name   string `json:"name"`
		Email  string `json:"email"`
	}

	type response struct {
		Type     string `json:"type"`
		Elements []Card `json:"elements"`
	}

	var result struct {
		Responses []response `json:"responses"`
	}

	w := c.Ctx.Response()

	w.Header().Set("Content-Type", "application/json")

	result.Responses = []response{
		{
			Type: "cards",
			Elements: []Card{
				{
					Title:    "Mario Chess",
					ImageURL: "https://image.ibb.co/hYWMXx/mariochess.jpg",
					Buttons: []Button{
						{
							Type:  Postback,
							Title: "Item Operations",
							Value: "fakeitemid1111",
						},
					},
				}, {
					Title:    "Nike Air",
					ImageURL: "https://image.ibb.co/iUGoCx/nike.jpg",
					Buttons: []Button{
						{
							Type:  Postback,
							Title: "Item Operations",
							Value: "fakeitemid2222",
						},
					},
				},
			},
		},
	}

	fmt.Printf("[JSON]: %+v.\n", result)

	json.NewEncoder(os.Stdout).Encode(result)

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		fmt.Printf("[JSON_ERR]: %s\n", err.Error())
	}

	fmt.Println("[DONE_JSON]")
	fmt.Printf("[RESP]: %+v\n", c.Ctx.Response())
	c.HTML(http.StatusOK)
}

// GetItemsAsCards gets the purchased or sold items for a user.
// Request Parameters:  user's first name
//						user's last name
//						user's email.
//						"purchased" or "sold".
// Return: A Cards object for the bot to display.
func (c *ChatWgtController) GetItemsAsCards() {
	fmt.Println("GetItemsAsCards!")
	req := c.Ctx.Request()
	err := req.ParseForm()
	if err != nil {
		fmt.Println("[ERR]: " + err.Error())
		c.HTML(http.StatusBadRequest)
		c.Ctx.Template = "error"
		return
	}
	//TODO: synchronous call to Mercari backend for the items of current user.

	// Return the list of items as Cards.
	type p struct {
		ItemID string `json:"itemid"`
		Name   string `json:"name"`
		Email  string `json:"email"`
	}

	type response struct {
		Type     string `json:"type"`
		Elements []Card `json:"elements"`
	}

	var result struct {
		Responses  []response `json:"responses"`
		Parameters p          `json:"parameters"`
	}

	w := c.Ctx.Response()

	w.Header().Set("Content-Type", "application/json")

	var buttons = []Button{
		{
			Type:  Goto,
			Title: "Where is my item?",
			Value: "5aa833dcf60bd80007b25375",
		}, {
			Type:  Goto,
			Title: "Request return",
			Value: "5aa83bbd7eefe000078cb066",
		}, {
			Type:  Goto,
			Title: "Cancel order",
			Value: "5aa83bcbf60bd80007b253e3",
		},
	}

	result.Responses = []response{
		{
			Type: "cards",
			Elements: []Card{
				{
					Title:    "Mario Chess",
					ImageURL: "https://image.ibb.co/hYWMXx/mariochess.jpg",
					Buttons:  buttons,
				}, {
					Title:    "Nike Air",
					ImageURL: "https://image.ibb.co/iUGoCx/nike.jpg",
					Buttons:  buttons,
				}, {
					Title:    "Nike Air",
					ImageURL: "https://image.ibb.co/iUGoCx/nike.jpg",
					Buttons:  buttons,
				},
			},
		},
	}

	result.Parameters.ItemID = "1234567890"

	fmt.Printf("[JSON]: %+v.\n", result)

	json.NewEncoder(os.Stdout).Encode(result)

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		fmt.Printf("[JSON_ERR]: %s\n", err.Error())
	}

	fmt.Println("[DONE_JSON]")
	fmt.Printf("[RESP]: %+v\n", c.Ctx.Response())
	c.HTML(http.StatusOK)

}

// ItemCurrentStatus is the webhook for getting current status of an item.
func (c *ChatWgtController) ItemCurrentStatus() {
	fmt.Println("ItemCurrentStatus!")
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
