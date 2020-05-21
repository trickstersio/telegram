package telegram

import "encoding/json"

const (
	DefaultURL = "https://api.telegram.org"

	MethodSendMessage = "sendMessage"
	MethodSetWebhook  = "setWebhook"

	CommandStart = "/start"
)

type Response struct {
	OK          bool            `json:"ok"`
	Description *string         `json:"description"`
	Result      json.RawMessage `json:"result"`
}

type Update struct {
	ID      int64    `json:"id"`
	Message *Message `json:"message"`
}

type Message struct {
	ID      int64    `json:"message_id"`
	Chat    Chat     `json:"chat"`
	Text    string   `json:"text"`
	From    *User    `json:"from"`
	Contact *Contact `json:"contact"`
}

type Contact struct {
	Phone  string `json:"phone_number"`
	UserID *int64 `json:"user_id"`
}

type User struct {
	ID int64 `json:"id"`
}

type Chat struct {
	ID int64 `json:"id"`
}

type SendMessageArgs struct {
	ChatID      int64       `json:"chat_id"`
	Text        string      `json:"text"`
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

type ReplyKeyboardMarkup struct {
	Keyboard        [][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool               `json:"resize_keyboard"`
	OneTimeKeyboard bool               `json:"one_time_keyboard"`
	Selective       bool               `json:"selective"`
}

type KeyboardButton struct {
	Text            string `json:"text"`
	RequestContact  bool   `json:"request_contact"`
	RequestLocation bool   `json:"request_location"`
}

type SetWebhookArgs struct {
	URL string `json:"url"`
}
