package config

import (
	"fmt"
	"log"
	"os"

	"github.com/cstockton/go-conv"
	env "github.com/joho/godotenv"
)

type Enviroment struct {
	SMTP     string
	Port     int
	Sender   string
	Password string
}

var Env Enviroment

func LoadEnv() {
	if err := env.Load(); err != nil {
		log.Fatal(err)
	}

	Env.SMTP = os.Getenv("EMAIL_SMTP")
	Env.Port, _ = conv.Int(os.Getenv("EMAIL_PORT"))
	Env.Sender = os.Getenv("EMAIL_SENDER")
	Env.Password = os.Getenv("EMAIL_PASSWORD")
	fmt.Println(Env)
}
