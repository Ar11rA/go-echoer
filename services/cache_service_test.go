package services

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockCacheClient struct {
	SetFunc func(ctx context.Context, key string, value string) error
	GetFunc func(ctx context.Context, key string) (string, error)
}

func (m *MockCacheClient) Set(ctx context.Context, key string, value string) error {
	if m.SetFunc != nil {
		return m.SetFunc(ctx, key, value)
	}
	return fmt.Errorf("Set is not defined")
}

func (m *MockCacheClient) Get(ctx context.Context, key string) (string, error) {
	if m.GetFunc != nil {
		return m.GetFunc(ctx, key)
	}
	return "", fmt.Errorf("Get is not defined")
}

func TestCacheServiceImpl_GetData(t *testing.T) {

	mockCacheClient := &MockCacheClient{
		GetFunc: func(ctx context.Context, key string) (string, error) {
			return "sample_value", nil
		},
	}

	cacheService := NewRedisService(mockCacheClient)

	// Act
	val, err := cacheService.GetData(context.Background(), "sample_key")

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "sample_value", val)
}

func TestCacheServiceImpl_SaveData(t *testing.T) {

	mockCacheClient := &MockCacheClient{
		SetFunc: func(ctx context.Context, key string, value string) error {
			return nil
		},
	}

	cacheService := NewRedisService(mockCacheClient)

	// Act
	err := cacheService.SaveData(context.Background(), "sample_key", "sample_value")

	// Assert
	assert.Nil(t, err)
}
