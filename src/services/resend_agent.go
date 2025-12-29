package services

import (
	"github.com/csvitor-dev/mail-sender/internal/config"
	"github.com/csvitor-dev/mail-sender/internal/entities"
	"github.com/resend/resend-go/v3"
)

func SendEmailByResend(job *entities.EmailJob) error {
	client := resend.NewClient(config.Env.API_KEY)
	params := &resend.SendEmailRequest{
		From:    config.Env.EMAIL_SENDER,
		To:      []string{job.To},
		Text:    job.Body,
		Subject: job.Subject,
	}
	var attach *resend.Attachment

	if job.HasFileToAttach() {
		content, _ := job.GetAttachment()
		attach = &resend.Attachment{
			Content:     content,
			ContentType: job.ContentType,
			Filename: job.Filename,
		}

		params.Attachments = []*resend.Attachment{attach}
	}
	_, err := client.Emails.Send(params)

	if err != nil {
		return err
	}
	return nil
}
