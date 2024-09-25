package services

import (
	"encoding/json"
	"fmt"
	"quote-server/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockHttpClient is a mock implementation of HttpClient for testing
type MockHttpClient struct {
	GetFunc  func(url string) ([]byte, error)
	PostFunc func(url string, body interface{}) ([]byte, error)
}

// Get mocks the Get method of HttpClient and returns raw byte data
func (m *MockHttpClient) Get(url string) ([]byte, error) {
	if m.GetFunc != nil {
		return m.GetFunc(url)
	}
	return nil, fmt.Errorf("GetFunc is not defined")
}

func (m *MockHttpClient) Post(url string, body interface{}) ([]byte, error) {
	if m.PostFunc != nil {
		return m.PostFunc(url, body)
	}
	return nil, fmt.Errorf("PostFunc is not defined")
}

func TestHttpServiceImpl_GetEcho(t *testing.T) {
	// Arrange
	mockResponse := types.EchoResponse{
		Args: map[string]interface{}{
			"query": "abc",
		},
	}

	// Mock the response as raw JSON bytes
	mockResponseBytes, _ := json.Marshal(mockResponse)

	mockHttpClient := &MockHttpClient{
		GetFunc: func(url string) ([]byte, error) {
			// Verify the correct URL is being passed
			expectedURL := "/get?query=abc"
			assert.Equal(t, expectedURL, url)

			// Return mock response bytes
			return mockResponseBytes, nil
		},
	}

	httpService := NewHttpService(mockHttpClient)

	// Act
	response, err := httpService.GetEcho("abc")

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "abc", response.Args["query"])
}

func TestHttpServiceImpl_GetEcho_Error(t *testing.T) {
	// Arrange
	mockHttpClient := &MockHttpClient{
		GetFunc: func(url string) ([]byte, error) {
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

func TestHttpServiceImpl_PostEcho(t *testing.T) {
	// Arrange
	mockRequest := types.EchoRequest{Text: "yo"}
	mockResponse := types.EchoResponse{
		Data: map[string]interface{}{
			"text": "yo",
		},
	}

	// Mock the response as raw JSON bytes
	mockResponseBytes, _ := json.Marshal(mockResponse)

	mockHttpClient := &MockHttpClient{
		PostFunc: func(url string, body interface{}) ([]byte, error) {
			expectedURL := "/post"
			assert.Equal(t, expectedURL, url)

			// Verify the request body
			expectedBody := types.EchoRequest{Text: "yo"}
			assert.Equal(t, expectedBody, body)

			// Return mock response bytes
			return mockResponseBytes, nil
		},
	}

	httpService := NewHttpService(mockHttpClient)

	// Act
	response, err := httpService.PostEcho(mockRequest)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "yo", response.Data["text"])
}

func TestHttpServiceImpl_PostEcho_Error(t *testing.T) {
	// Arrange
	mockHttpClient := &MockHttpClient{
		PostFunc: func(url string, body interface{}) ([]byte, error) {
			return nil, fmt.Errorf("network error")
		},
	}

	httpService := NewHttpService(mockHttpClient)

	// Act
	response, err := httpService.PostEcho(types.EchoRequest{Text: "yo"})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Contains(t, err.Error(), "network error")
}

// TestHttpServiceImpl_GetQuotes tests the GetQuotes function
func TestHttpServiceImpl_GetQuotes(t *testing.T) {
	// Arrange
	limit := int32(2)
	mockQuote := &types.QuoteResponse{Content: "This is a quote!"}
	mockResponseBytes, _ := json.Marshal(mockQuote)

	mockHttpClient := &MockHttpClient{
		GetFunc: func(url string) ([]byte, error) {
			return mockResponseBytes, nil // Simulate a successful response
		},
	}

	httpService := NewHttpService(mockHttpClient)

	// Act
	quotes, err := httpService.GetQuotes(limit)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, quotes)
	assert.Len(t, quotes, int(limit))
	assert.Equal(t, mockQuote.Content, quotes[0].Content)
}

func TestHttpServiceImpl_GetQuotes_Error(t *testing.T) {
	// Arrange
	limit := int32(2)

	mockHttpClient := &MockHttpClient{
		GetFunc: func(url string) ([]byte, error) {
			return nil, fmt.Errorf("network error") // Simulate an error response
		},
	}

	httpService := NewHttpService(mockHttpClient)

	// Act
	quotes, err := httpService.GetQuotes(limit)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, quotes)
	assert.Contains(t, err.Error(), "network error")
}
