package telegram

import (
	"testing"

	"github.com/matryer/is"
)

func TestClient_SetWebhook(t *testing.T) {
	assert := is.New(t)

	client := NewClient("test", WithMethodMock(MethodSetWebhook, true, nil))

	err := client.SetWebhook(SetWebhookArgs{})

	assert.NoErr(err)
}

func TestClient_SendMessage(t *testing.T) {
	assert := is.New(t)

	mock := Message{
		Text: "Hello from Telegram",
	}

	client := NewClient("test", WithMethodMock(MethodSendMessage, true, mock))

	msg, err := client.SendMessage(SendMessageArgs{})

	assert.NoErr(err)
	assert.Equal(mock.Text, msg.Text)
}
