package entities

type EmailJob struct {
	To          string
	Subject     string
	Body        string
	ContentType string
	FilePath    string
}

func NewEmail(to, subject, body, contentType, path string) *EmailJob {
	return &EmailJob{
		To:          to,
		Subject:     subject,
		Body:        body,
		ContentType: contentType,
		FilePath:    path,
	}
}

func (e *EmailJob) HasFileToAttach() bool {
	return e.FilePath != ""
}
