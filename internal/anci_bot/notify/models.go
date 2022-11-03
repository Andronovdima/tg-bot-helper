package notify

// Create a struct that mimics the webhook response body
// https://core.telegram.org/bots/api#update

type Message struct {
	Text string `json:"text"`
	Chat struct {
		ID int64 `json:"id"`
	} `json:"chat"`
}

type notifyRequest struct {
	Message Message `json:"message"`
}
