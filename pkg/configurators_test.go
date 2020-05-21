package telegram

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"
)

func TestWithHTTPClient(t *testing.T) {
	assert := is.New(t)

	in := Message{
		Text: "Hello from Telegram!",
	}

	client := NewClient("test", WithClientMock(true, in))

	out, err := client.SendMessage(SendMessageArgs{})

	assert.NoErr(err)
	assert.Equal(out.Text, in.Text)
}

func TestWithURL(t *testing.T) {
	assert := is.New(t)

	in := Message{
		Text: "Hello from Telegram!",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(in)

		assert.NoErr(err)

		err = json.NewEncoder(w).Encode(Response{
			OK:     true,
			Result: data,
		})

		assert.NoErr(err)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	client := NewClient("test", WithURL(server.URL))

	out, err := client.SendMessage(SendMessageArgs{})

	assert.NoErr(err)
	assert.Equal(out.Text, in.Text)
}
