package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"quote-server/types"
)

type HttpClient interface {
	Get(url string) (*types.EchoResponse, error)
}

type EchoHttpClient struct{}

func (c *EchoHttpClient) Get(url string) (*types.EchoResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response
	var postmanResponse types.EchoResponse
	if err := json.Unmarshal(body, &postmanResponse); err != nil {
		return nil, err
	}

	return &postmanResponse, nil
}
