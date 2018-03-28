// Package models has the data models that are common among all inputs that BotEngine send to our webhooks.
package models

// Interaction is for the "interaction" object in webhooks' input from BotEngine.
type Interaction struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Action string `json:"action"`
}

// Parameters is for the "parameters" object in webhooks' input from BotEngine.
type Parameters struct {
	DefaultURL string `json:"default_url"`
	Any        string `json:"any"`
}

// Fulfillment is for "fulfillment".
type Fulfillment struct {
	Type      string `json:"text"`
	Delay     int    `json:"delay, omitempty"`
	Message   string `json:"message, omitempty"`
	WebhookID string `json:"webhookId, omitempty"`
}

// Status is for the "status" object in what BotEngine feeds our webhooks.
type Status struct {
	Code int    `json:"code"`
	Type string `json:"type"`
}
