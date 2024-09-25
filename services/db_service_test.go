package services

import (
	"fmt"
	"quote-server/types"
	"quote-server/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockRow struct {
	ID       int64
	Username string
	Email    string
}

func (m *MockRow) Scan(dest ...interface{}) error {
	if len(dest) < 3 {
		return fmt.Errorf("not enough destinations")
	}
	*dest[0].(*int64) = m.ID
	*dest[1].(*string) = m.Username
	*dest[2].(*string) = m.Email
	return nil
}

type MockDBClient struct {
	ExecFunc  func(query string, args ...interface{}) error
	QueryFunc func(query string, args ...interface{}) utils.Row
}

func (m *MockDBClient) Exec(query string, args ...interface{}) error {
	if m.ExecFunc != nil {
		return m.ExecFunc(query, "")
	}
	return fmt.Errorf("Exec is not defined")
}

func (m *MockDBClient) Query(query string, args ...interface{}) utils.Row {
	if m.QueryFunc != nil {
		return m.QueryFunc(query, "")
	}
	return nil
}

func TestDbServiceImpl_InsertUser(t *testing.T) {

	mockDbClient := &MockDBClient{
		ExecFunc: func(query string, args ...interface{}) error {
			return nil
		},
	}

	dbService := NewDbService(mockDbClient)

	// Act
	err := dbService.InsertUser(types.UserModel{})

	// Assert
	assert.Nil(t, err)
}

func TestDbServiceImpl_GetUser(t *testing.T) {

	mockDbClient := &MockDBClient{
		QueryFunc: func(query string, args ...interface{}) utils.Row {
			return &MockRow{ID: 1, Username: "testuser", Email: "test@example.com"}
		},
	}

	dbService := NewDbService(mockDbClient)

	// Act
	resp, err := dbService.GetUserById(int32(5))

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "test@example.com", resp.Email)
}
