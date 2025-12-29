package entities

import (
	"os"
	"path/filepath"
	"strings"
)

type EmailJob struct {
	To          string
	Subject     string
	Body        string
	ContentType string
	FilePath    string
	Filename    string
}

func NewEmail(to, subject, body, contentType, path string) (*EmailJob, error) {
	absolutePath, err := filepath.Abs(path)

	if err != nil {
		return nil, err
	}

	return &EmailJob{
		To:          to,
		Subject:     subject,
		Body:        body,
		ContentType: contentType,
		FilePath:    absolutePath,
		Filename: getFileName(absolutePath),
	}, nil
}

func getFileName(path string) string {
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
}

func (e *EmailJob) HasFileToAttach() bool {
	return e.FilePath != ""
}

func (e *EmailJob) GetAttachment() ([]byte, error) {
	payload, err := os.ReadFile(e.FilePath)

	if err != nil {
		return nil, err
	}
	return payload, nil
}
