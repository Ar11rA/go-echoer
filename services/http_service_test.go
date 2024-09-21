package services

import (
	"fmt"
	"quote-server/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockHttpClient is a mock implementation of HttpClient for testing
type MockHttpClient struct {
	GetFunc func(url string) (*types.EchoResponse, error)
}

// Get mocks the Get method of HttpClient
func (m *MockHttpClient) Get(url string) (*types.EchoResponse, error) {
	if m.GetFunc != nil {
		return m.GetFunc(url)
	}
	return nil, fmt.Errorf("GetFunc is not defined")
}

func TestHttpServiceImpl_GetEcho(t *testing.T) {
	// Arrange
	mockHttpClient := &MockHttpClient{
		GetFunc: func(url string) (*types.EchoResponse, error) {
			// Verify the correct URL is being passed
			expectedURL := "https://postman-echo.com/get?query=abc"
			assert.Equal(t, expectedURL, url)

			// Return a mock response
			return &types.EchoResponse{
				Args: struct {
					Query string `json:"query"`
				}{
					Query: "abc",
				},
			}, nil
		},
	}

	httpService := NewHttpService(mockHttpClient)

	// Act
	response, err := httpService.GetEcho("abc")

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "abc", response.Args.Query)
}

func TestHttpServiceImpl_GetEcho_Error(t *testing.T) {
	// Arrange
	mockHttpClient := &MockHttpClient{
		GetFunc: func(url string) (*types.EchoResponse, error) {
			return nil, fmt.Errorf("network error")
		},
	}

	httpService := NewHttpService(mockHttpClient)

	// Act
	response, err := httpService.GetEcho("abc")

	// Assert
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Contains(t, err.Error(), "network error")
}
