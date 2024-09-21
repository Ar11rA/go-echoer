package services

import (
	"path/filepath"
	"quote-server/utils"
	"time"
)

// FileService defines the interface for file operations
type FileService interface {
	Save(content string) error
}

// FileServiceImpl is an implementation of FileService
type FileServiceImpl struct {
	Directory  string
	FileWriter utils.FileWriter
}

// Save writes the content to a timestamped file
func (fs *FileServiceImpl) Save(content string) error {
	filename := time.Now().Format("2006-01-02T15:04:05.000") + ".txt"
	filepath := filepath.Join(fs.Directory, filename)
	return fs.FileWriter.WriteFile(filepath, []byte(content), 0644)
}
