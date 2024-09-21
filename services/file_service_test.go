package services

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockFileWriter is a mock implementation of FileWriter for testing
type MockFileWriter struct {
	WriteFileFunc func(filename string, content []byte, perm os.FileMode) error
}

// WriteFile mocks the WriteFile method
func (m *MockFileWriter) WriteFile(filename string, content []byte, perm os.FileMode) error {
	if m.WriteFileFunc != nil {
		return m.WriteFileFunc(filename, content, perm)
	}
	return nil
}

func TestFileServiceImpl_Save(t *testing.T) {
	// Arrange
	mockFileWriter := &MockFileWriter{
		WriteFileFunc: func(filename string, content []byte, perm os.FileMode) error {
			// Verify the filename and content
			assert.Contains(t, filename, "txt") // Adjust time format if necessary
			assert.Equal(t, []byte("test content"), content)
			assert.Equal(t, os.FileMode(0644), perm)
			return nil // Simulate successful file write
		},
	}

	fileService := &FileServiceImpl{
		Directory:  "./",
		FileWriter: mockFileWriter,
	}

	// Act
	err := fileService.Save("test content")

	// Assert
	assert.NoError(t, err)
}
