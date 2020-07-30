package telegram

import (
	"net/http"
)

type Configurator interface {
	Configure(client *Client)
}

type ConfiguratorFunc func(client *Client)

func (f ConfiguratorFunc) Configure(client *Client) {
	f(client)
}

func WithHTTPClient(httpClient *http.Client) ConfiguratorFunc {
	return func(client *Client) {
		client.httpClient = httpClient
	}
}

func WithURL(url string) ConfiguratorFunc {
	return func(client *Client) {
		client.url = url
	}
}
