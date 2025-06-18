package services

import (
	"crypto/tls"
	"fmt"

	e "github.com/csvitor-dev/mail-sender/internal/entities"
	"github.com/csvitor-dev/mail-sender/pkg/config"
	"gopkg.in/mail.v2"
)

func SendEmail(job e.EmailJob, smtp config.SMTPConfig) error {
	mailer := mail.NewMessage()

	mailer.SetHeaders(map[string][]string{
		"From":    {smtp.Username},
		"To":      {job.To},
		"Subject": {job.Subject},
	})
	mailer.SetBody("text/plain", job.Body)
	mailer.Attach(job.FilePath)

	dialer := createMailDialer(smtp)

	try := 2
	var err error

	for try > 0 {
		if err = dialer.DialAndSend(mailer); err == nil {
			break;
		}
		try--
	}

	if err != nil {
		return fmt.Errorf("fail to send email: %w", err)
	}
	return nil
}

func createMailDialer(smtp config.SMTPConfig) *mail.Dialer {
	dialer := mail.NewDialer(smtp.Host, smtp.Port, smtp.Username, smtp.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return dialer
}
