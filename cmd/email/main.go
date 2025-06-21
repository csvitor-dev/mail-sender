package main

import (
	"log"

	"github.com/csvitor-dev/mail-sender/internal/entities"
	"github.com/csvitor-dev/mail-sender/internal/services"
	"github.com/csvitor-dev/mail-sender/pkg/config"
	src_service "github.com/csvitor-dev/mail-sender/src/services"
)

func main() {
	smtpConfig := config.NewSMTP()
	numWorkers := config.Env.WORKER_NUM

	jobs := make(chan *entities.EmailJob, 100)
	wg := src_service.StartWorkerPool(numWorkers, jobs, smtpConfig)

	services.LoadFiles(jobs)
	close(jobs)
	wg.Wait()

	log.Println("All emails have been processed!")
	log.Printf("Total Requests: %d\n", src_service.TotalRequests)
	log.Printf("Fail email requests: %v\n", src_service.FailRequests)
	log.Printf("Success email requests: %v\n", src_service.SuccessRequests)
}
