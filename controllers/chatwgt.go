package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gernest/utron/controller"
	"github.com/nikkitan/livechat-demo/controllers/models"
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
			"get;/chatwgt/getsingleitem;VerifyWebhook",
			"post;/chatwgt/getsingleitem;GetSelectedItemCard",
		},
	}
}

//Home is responsible for rendering home page.
func (c *ChatWgtController) Home() {
	c.Ctx.Template = "chatwgt_home"
	c.HTML(http.StatusOK)
}

//Test is responsible for rendering home page.
func (c *ChatWgtController) Test() {
	fmt.Println("Test!")
	c.HTML(http.StatusOK)
}

//GetItemStatus is responsible for rendering home page.
func (c *ChatWgtController) GetItemStatus() {
	fmt.Println("GetItemStatus!")

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
	type p struct {
		ItemID     string `json:"current_item_id"`
		ItemImgURL string `json:"selected_img_url"`
		ItemName   string `json:"curr_item_name"`
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

	result.Parameters.ItemID = "fakeitem1111"
	result.Parameters.ItemImgURL = "https://image.ibb.co/hYWMXx/mariochess.jpg"
	result.Parameters.ItemName = "Mario Chess"

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
	fmt.Printf("[REQ_BODY]: %+v\n", req.Body)

	err = req.ParseForm()

	if err != nil {
		fmt.Println("[ERR]: " + err.Error())
		c.HTML(http.StatusBadRequest)
		c.Ctx.Template = "error"
		return
	}

	// Generate the Quick Replies.
	type p struct {
		ItemID string `json:"current_item_id"`
	}

	type response struct {
		Type    string          `json:"type"`
		Title   string          `json:"title"`
		Buttons []models.Button `json:"buttons"`
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
			Buttons: []models.Button{
				{
					Type:  models.Goto,
					Title: "Item Status",
					Value: "5aa833dcf60bd80007b25375",
				}, {
					Type:  models.Goto,
					Title: "Cancel Item",
					Value: "5aa83bcbf60bd80007b253e3",
				}, {
					Type:  models.Goto,
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
		Type     string        `json:"type"`
		Elements []models.Card `json:"elements"`
	}

	var result struct {
		Responses []response `json:"responses"`
	}

	w := c.Ctx.Response()

	w.Header().Set("Content-Type", "application/json")

	result.Responses = []response{
		{
			Type: "cards",
			Elements: []models.Card{
				{
					Title:    "Mario Chess",
					ImageURL: "https://image.ibb.co/hYWMXx/mariochess.jpg",
					Buttons: []models.Button{
						{
							Type:  models.Postback,
							Title: "Item Operations",
							Value: "fakeitemid1111",
						}, {
							Type:  models.Postback,
							Title: "Mario Chess",
							Value: "Item Name%Mario Chess",
						},
					},
				}, {
					Title:    "Nike Air",
					ImageURL: "https://image.ibb.co/iUGoCx/nike.jpg",
					Buttons: []models.Button{
						{
							Type:  models.Postback,
							Title: "Item Operations",
							Value: "fakeitemid2222",
						}, {
							Type:  models.Postback,
							Title: "Nike Air",
							Value: "Item Name%Nike Air",
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
		Type     string        `json:"type"`
		Elements []models.Card `json:"elements"`
	}

	var result struct {
		Responses []response `json:"responses"`
	}

	w := c.Ctx.Response()

	w.Header().Set("Content-Type", "application/json")

	result.Responses = []response{
		{
			Type: "cards",
			Elements: []models.Card{
				{
					Title:    "Mario Chess",
					ImageURL: "https://image.ibb.co/hYWMXx/mariochess.jpg",
					Buttons: []models.Button{
						{
							Type:  models.Postback,
							Title: "Item Operations",
							Value: "fakeitemid1111",
						}, {
							Type:  models.Postback,
							Title: "Mario Chess",
							Value: "Item Name%Mario Chess",
						},
					},
				}, {
					Title:    "Nike Air",
					ImageURL: "https://image.ibb.co/iUGoCx/nike.jpg",
					Buttons: []models.Button{
						{
							Type:  models.Postback,
							Title: "Item Operations",
							Value: "fakeitemid2222",
						}, {
							Type:  models.Postback,
							Title: "Nike Air",
							Value: "Item Name%Nike Air",
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

// GetSelectedItemCard gets the selected item's info from backend
// and generates a Card object for it.
// Request Parameters:  item ID
// Return: A Cards object for the bot GetSelectedItemCardto display.
func (c *ChatWgtController) GetSelectedItemCard() {
	fmt.Println("GetSelectedItemCard!")
	req := c.Ctx.Request()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("Error reading body: %v", err)
		return
	}
	//TODO: Parse the JSON in body.
	type contextParameters struct {
		DefaultURL       string `json:"default_url"`
		Any              string `json:"any, omitempty"`
		DefaultEMail     string `json:"default_email"`
		DefaultFirstname string `json:"default_username"`
		DefaultLastname  string `json:"default_lastname"`
	}

	type context struct {
		ID         string            `json:"id, omitempty"`
		Name       string            `json:"name, omitempty"`
		Parameters contextParameters `json:"parameters"`
	}

	type result struct {
		Source        string               `json:"source"`
		ResolvedQuery string               `json:"resolvedQuery"`
		Goto          string               `json:"goto"`
		Confidence    int                  `json:"confidence"`
		Score         int                  `json:"score"`
		LifeSpan      int                  `json:"lifespan"`
		Incomplete    bool                 `json:"incomplete"`
		StoryID       string               `json:"storeId"`
		Interaction   models.Interaction   `json:"interaction"`
		Parameters    models.Parameters    `json:"parameters"`
		Contexts      []context            `json:"contexts"`
		Fulfillment   []models.Fulfillment `json:"fulfillment"`
	}

	var webhookInput struct {
		Timestamp string        `json:"timestamp"`
		SessionID string        `json:"sessionId"`
		Result    result        `json:"result"`
		Status    models.Status `json:"status"`
	}

	if err = json.Unmarshal(body, &webhookInput); err != nil {
		fmt.Printf("[JSON_ERR]: %s.\n", err.Error())
		c.HTML(http.StatusInternalServerError)
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

	// TODO: get item's data from backend by passing item ID to the backend.

	//Generate the Quick Replies.
	type p struct {
		ItemID string `json:"current_item_id"`
	}

	// A Card object.
	type response struct {
		Type     string          `json:"type"`
		Title    string          `json:"title"`
		Subtitle string          `json:"subtitle"`
		ImageURL string          `json:"imageUrl"`
		Buttons  []models.Button `json:"buttons"`
	}

	var card struct {
		Responses []response `json:"responses"`
		//Parameters p          `json:"parameters"`
	}

	w := c.Ctx.Response()

	w.Header().Set("Content-Type", "application/json")

	card.Responses = []response{
		{
			Type:  "card",
			Title: "Mario Chess",

			Buttons: []models.Button{
				{
					Type:  models.Goto,
					Title: "Item Status",
					Value: "5aa833dcf60bd80007b25375",
				}, {
					Type:  models.Goto,
					Title: "Cancel Item",
					Value: "5aa83bcbf60bd80007b253e3",
				}, {
					Type:  models.Goto,
					Title: "Return Item",
					Value: "5aa83bbd7eefe000078cb066",
				},
			},
		},
	}

	fmt.Printf("[JSON]: %+v.\n", card)

	json.NewEncoder(os.Stdout).Encode(card)

	err = json.NewEncoder(w).Encode(card)
	if err != nil {
		fmt.Printf("[JSON_ERR]: %s\n", err.Error())
	}

	fmt.Println("[DONE_JSON]")
	fmt.Printf("[RESP]: %+v\n", c.Ctx.Response())

	// Synchronous POST to trigger
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
