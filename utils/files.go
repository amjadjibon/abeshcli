package utils

import (
	"fmt"
	"os"
)

// exists reports whether the named file or directory exists
func exists(path string, isDir bool) bool {
	if path == "" {
		fmt.Println("path is empty")
		return false
	}

	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) { // look for the specific error type
			return false
		}
	}

	return isDir == info.IsDir()
}

// FolderExists reports whether the provided directory exists.
func FolderExists(path string) bool {
	return exists(path, true)
}

// FileExists reports whether the provided directory exists.
func FileExists(path string) bool {
	return exists(path, false)
}
