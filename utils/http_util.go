package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type HttpClient interface {
	Get(url string) ([]byte, error)
	Post(url string, body interface{}) ([]byte, error)
}

// EchoHttpClient implements HttpClient and allows generic parsing
type EchoHttpClient struct{}

// Get fetches the URL and returns the raw response body as bytes
func (c *EchoHttpClient) Get(url string) ([]byte, error) {
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

	return body, nil
}

func (c *EchoHttpClient) Post(url string, body interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

// ParseResponse is a generic method that takes the raw response body and unmarshals it into a provided type T
func ParseResponse[T any](data []byte) (*T, error) {
	var result T
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
