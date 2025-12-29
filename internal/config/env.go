package config

import (
	"fmt"
	"log"
	"os"

	"github.com/cstockton/go-conv"
	env "github.com/joho/godotenv"
)

type Enviroment struct {
	API_KEY        string
	EMAIL_SMTP     string
	EMAIL_PORT     int
	EMAIL_SENDER   string
	EMAIL_USER     string
	EMAIL_PASSWORD string
	WORKER_NUM     int
}

var Env Enviroment

func LoadEnv() {
	if err := env.Load(); err != nil {
		log.Fatal(err)
	}

	Env.API_KEY = os.Getenv("API_KEY")
	Env.EMAIL_SMTP = os.Getenv("EMAIL_SMTP")
	Env.EMAIL_PORT, _ = conv.Int(os.Getenv("EMAIL_PORT"))
	Env.EMAIL_SENDER = os.Getenv("EMAIL_SENDER")
	Env.EMAIL_USER = os.Getenv("EMAIL_USER")
	Env.EMAIL_PASSWORD = os.Getenv("EMAIL_PASSWORD")
	Env.WORKER_NUM, _ = conv.Int(os.Getenv("WORKER_NUM"))

	fmt.Print("The environment was loaded successfully!\n\n")
}
