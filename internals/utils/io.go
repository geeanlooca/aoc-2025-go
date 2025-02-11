package utils

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kardianos/osext"
)

// FileExists checks whether a file exists
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return os.IsExist(err)
}

// ReadFileLines reads lines from a text file
func ReadFileLines(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening the file %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error when reading file: %w", err)
	}

	return lines, nil
}

// GetExecutablePath gets the path of the executable
func GetExecutablePath() (string, error) {
	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		return "", fmt.Errorf("could not get the executable path")
	}

	return folderPath, nil
}
