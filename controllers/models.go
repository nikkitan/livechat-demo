package controllers

// Card is for when we want to return a "Cards" object to be displayed in the bot.
type Card struct {
	Title    string   `json:"title"`
	Subtitle string   `json:"subtitle, omitempty"`
	ImageURL string   `json:"imageUrl, omitempty"`
	Buttons  []Button `json:"buttons"`
	// TODO: Define for "Filters" object, which is optional and undocumented on BotEngine's website.
}

var (
	// Goto can be used for things like Button type.
	Goto = "goto"
	// Postback can be used for things like Button type.
	Postback = "postback"
	// URL can be used for things like Button type.
	URL = "url"
	// Phone can be used for things like Button type.
	Phone = "phone"
)

// Button is for when we want to return a "button" object to be displayed in the bot.
type Button struct {
	Type  string `json:"type"`
	Title string `json:"title"`
	Value string `json:"value"`
}
