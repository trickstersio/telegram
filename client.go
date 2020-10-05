package telegram

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Client struct {
	token      string
	url        string
	httpClient *http.Client
}

func NewClient(token string, configurators ...Configurator) *Client {
	client := &Client{
		token:      token,
		httpClient: http.DefaultClient,
		url:        DefaultURL,
	}

	for _, configurator := range configurators {
		configurator.Configure(client)
	}

	return client
}

func (t *Client) Call(ctx context.Context, method string, args interface{}, out interface{}) error {
	httpRequestBody, err := json.Marshal(args)
	methodURL := fmt.Sprintf("%s/bot%s/%s", t.url, t.token, method)

	if err != nil {
		return fmt.Errorf("failed to serialize arguments: %w", err)
	}

	httpRequest, err := http.NewRequest(
		http.MethodPost,
		methodURL,
		bytes.NewReader(httpRequestBody),
	)

	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	httpRequest.Header.Set("Accept", "application/json")
	httpRequest.Header.Set("Content-Type", "application/json")

	httpResponse, err := t.httpClient.Do(httpRequest.WithContext(ctx))

	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	defer func() {
		if err := httpResponse.Body.Close(); err != nil {
			log.Println("Failed to close telegram response body", err)
		}
	}()

	var telegramResponse Response

	if err := json.NewDecoder(httpResponse.Body).Decode(&telegramResponse); err != nil {
		return fmt.Errorf("failed to decode telegram response body: %w", err)
	}

	if !telegramResponse.OK {
		var description string

		if telegramResponse.Description != nil {
			description = *telegramResponse.Description
		}

		return fmt.Errorf("response is not ok: %s", description)
	}

	if out != nil {
		if err := json.Unmarshal(telegramResponse.Result, out); err != nil {
			return fmt.Errorf("failed to decode telegram response result: %w", err)
		}
	}

	return nil
}

func (t *Client) SetWebhook(ctx context.Context, args SetWebhookArgs) error {
	if err := t.Call(ctx, MethodSetWebhook, args, nil); err != nil {
		return fmt.Errorf("failed to call method %s: %w", MethodSetWebhook, err)
	}

	return nil
}

func (t *Client) SendMessage(ctx context.Context, args SendMessageArgs) (Message, error) {
	var message Message

	if err := t.Call(ctx, MethodSendMessage, args, &message); err != nil {
		return Message{}, fmt.Errorf("failed to call method %s: %w", MethodSendMessage, err)
	}

	return message, nil
}
