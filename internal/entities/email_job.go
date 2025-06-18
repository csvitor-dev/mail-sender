package entities

type EmailJob struct {
	To       string
	Subject  string
	Body     string
	FilePath string
}

func NewEmail(to, subject, body, path string) EmailJob {
	return EmailJob{
		To:       to,
		Subject:  subject,
		Body:     body,
		FilePath: path,
	}
}
