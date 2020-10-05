package telegram

import (
	"context"
	"testing"

	"github.com/matryer/is"
)

func TestClient_SetWebhook(t *testing.T) {
	assert := is.New(t)
	ctx := context.Background()

	client := NewClient("test", WithMethodMock(MethodSetWebhook, true, nil))

	err := client.SetWebhook(ctx, SetWebhookArgs{})

	assert.NoErr(err)
}

func TestClient_SendMessage(t *testing.T) {
	assert := is.New(t)
	ctx := context.Background()

	mock := Message{
		Text: "Hello from Telegram",
	}

	client := NewClient("test", WithMethodMock(MethodSendMessage, true, mock))

	msg, err := client.SendMessage(ctx, SendMessageArgs{})

	assert.NoErr(err)
	assert.Equal(mock.Text, msg.Text)
}
