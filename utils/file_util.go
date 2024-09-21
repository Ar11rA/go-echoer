package utils

import "os"

// FileWriter defines an interface for writing content to a file
type FileWriter interface {
	WriteFile(filename string, content []byte, perm os.FileMode) error
}

// OSFileWriter is an implementation of FileWriter that uses os.WriteFile
type OSFileWriter struct{}

// WriteFile writes content to a file using os.WriteFile
func (w *OSFileWriter) WriteFile(filename string, content []byte, perm os.FileMode) error {
	return os.WriteFile(filename, content, perm)
}
