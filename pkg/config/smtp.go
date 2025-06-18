package config

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

func NewSMTP() SMTPConfig {
	LoadEnv()

	return SMTPConfig{
		Host:     Env.SMTP,
		Port:     Env.Port,
		Username: Env.Sender,
		Password: Env.Password,
	}
}
