package services

import (
	"crypto/tls"
	"fmt"

	e "github.com/csvitor-dev/mail-sender/internal/entities"
	"github.com/csvitor-dev/mail-sender/pkg/config"
	"gopkg.in/mail.v2"
)

func SendEmail(job *e.EmailJob, smtp config.SMTPConfig) error {
	mailer := mail.NewMessage()

	mailer.SetHeaders(map[string][]string{
		"From":    {mailer.FormatAddress(smtp.Email, smtp.Username)},
		"To":      {job.To},
		"Subject": {job.Subject},
	})
	mailer.SetBody(job.ContentType, job.Body)

	if job.HasFileToAttach() {
		mailer.Attach(job.FilePath)
	}
	dialer := createMailDialer(smtp)

	if err := dialer.DialAndSend(mailer); err == nil {
		return fmt.Errorf("fail to send email: %w", err)
	}
	return nil
}

func createMailDialer(options config.SMTPConfig) *mail.Dialer {
	dialer := mail.NewDialer(options.Host, options.Port, options.Email, options.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return dialer
}
