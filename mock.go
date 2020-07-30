package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
)

func WithClientMock(ok bool, result interface{}) ConfiguratorFunc {
	return WithHTTPClient(&http.Client{
		Transport: &MethodMock{
			ok:     ok,
			result: result,
		},
	})
}

func WithMethodMock(name string, ok bool, result interface{}) ConfiguratorFunc {
	return WithHTTPClient(&http.Client{
		Transport: &MethodMock{
			name:   &name,
			ok:     ok,
			result: result,
		},
	})
}

type MethodMock struct {
	ok     bool
	result interface{}
	name   *string
}

func (mock *MethodMock) RoundTrip(r *http.Request) (*http.Response, error) {
	if mock.name != nil {
		if name := path.Base(r.URL.Path); name != *mock.name {
			return nil, fmt.Errorf("invalid method call: called %s when expected %s", name, *mock.name)
		}
	}

	resultData, err := json.Marshal(mock.result)

	if err != nil {
		return nil, fmt.Errorf("failed to mock telegram request: %w", err)
	}

	responseData, err := json.Marshal(Response{
		OK:     mock.ok,
		Result: resultData,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to mock telegram request: %w", err)
	}

	return &http.Response{
		Status:     http.StatusText(http.StatusOK),
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewReader(responseData)),
	}, nil
}
