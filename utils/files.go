package utils

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
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

func File2lines(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return LinesFromReader(f)
}

func LinesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// InsertStringToFile ...
func InsertStringToFile(path, str string, index int) error {
	lines, err := File2lines(path)
	if err != nil {
		return err
	}

	fileContent := ""
	for i, line := range lines {
		if i == index {
			fileContent += str
			fileContent += "\n"
		}
		fileContent += line
		fileContent += "\n"
	}

	return ioutil.WriteFile(path, []byte(fileContent), 0644)
}

func InsertStringToFile2(path, str string) error {
	lines, err := File2lines(path)
	if err != nil {
		return err
	}

	fileContent := ""
	for _, line := range lines {
		if line == ")" {
			fileContent += str
			fileContent += "\n"
		}
		fileContent += line
		fileContent += "\n"
	}

	return ioutil.WriteFile(path, []byte(fileContent), 0644)
}
