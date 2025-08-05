package config

type SMTPConfig struct {
	Host     string
	Port     int
	Email    string
	Username string
	Password string
}

func NewSMTP() SMTPConfig {
	LoadEnv()

	return SMTPConfig{
		Host:     Env.EMAIL_SMTP,
		Port:     Env.EMAIL_PORT,
		Email:    Env.EMAIL_SENDER,
		Username: Env.EMAIL_USER,
		Password: Env.EMAIL_PASSWORD,
	}
}
