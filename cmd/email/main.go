package main

import (
	"log"
	"strings"

	"github.com/csvitor-dev/mail-sender/internal/entities"
	"github.com/csvitor-dev/mail-sender/internal/services/files"
	"github.com/csvitor-dev/mail-sender/pkg/cli"
	"github.com/csvitor-dev/mail-sender/pkg/config"
	src_service "github.com/csvitor-dev/mail-sender/src/services"
)

func main() {
	smtpConfig := config.NewSMTP()
	numWorkers := config.Env.WORKER_NUM

	courseAndProofNum := strings.ToLower(cli.Prompt("Enter the target course (CC or ES) with number (1, 2, 3)"))

	if courseAndProofNum[:2] != "cc" && courseAndProofNum[:2] != "es" {
		log.Fatalf("Invalid course: %s. Please enter 'CC' or 'ES'.", courseAndProofNum)
	}
	subject := cli.Prompt("Enter the subject for the emails")
	bodyMessage := cli.Prompt("Enter the body message for the emails")

	jobs := make(chan *entities.EmailJob, 100)
	wg := src_service.StartWorkerPool(numWorkers, jobs, smtpConfig)

	files.LoadFromFiles(courseAndProofNum, subject, bodyMessage, jobs)
	close(jobs)
	wg.Wait()

	log.Println("All emails have been processed!")
	log.Printf("Total Requests: %d\n", src_service.TotalRequests)
	log.Printf("Fail email requests: %v\n", src_service.FailRequests)
	log.Printf("Success email requests: %v\n", src_service.SuccessRequests)
}
