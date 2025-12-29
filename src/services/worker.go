package services

import (
	"log"
	"sync"
	"time"

	"github.com/csvitor-dev/mail-sender/internal/entities"
)

var (
	SuccessRequests = []string{}
	FailRequests    = []string{}
	TotalRequests   = 0
)

func StartWorkerPool(workers int, jobs <-chan *entities.EmailJob) *sync.WaitGroup {
	var group sync.WaitGroup

	for i := 1; i <= workers; i++ {
		group.Add(1)
		go worker(i, jobs, &group)
	}
	return &group
}

func worker(id int, jobs <-chan *entities.EmailJob, group *sync.WaitGroup) {
	defer group.Done()

	for job := range jobs {
		log.Printf("Worker [%d]: Job processing for %v\n", id, job.To)

		if err := SendEmailByResend(job); err != nil {
			log.Printf("Worker [%d]: Fail to send email for %v: %v", id, job.To, err)
			FailRequests = append(FailRequests, job.To)
		} else {
			log.Printf("Worker [%d]: Email submitted successfully %v", id, job.To)
			SuccessRequests = append(SuccessRequests, job.To)
		}
		TotalRequests++
		time.Sleep(time.Second * 5)
	}
}
