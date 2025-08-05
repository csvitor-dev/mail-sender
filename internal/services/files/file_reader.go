package files

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func GetAbsolutePath(target, fileName string) (string, error) {
	filePath := target + fileName
	absolutePath, err := filepath.Abs(filePath)

	if err != nil {
		return "", err
	}
	return absolutePath, nil
}

func ReadFile(fileName string) ([][]string, error) {
	path, err := GetAbsolutePath("internal/emails/", fileName)

	if err != nil {
		return nil, err
	}
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var content [][]string
	for scanner.Scan() {
		emailAndFile := strings.Split(scanner.Text(), ",")
		content = append(content, emailAndFile)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return content, nil
}
