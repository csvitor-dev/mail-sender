package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/csvitor-dev/mail-sender/internal/services/files"
	"github.com/csvitor-dev/mail-sender/pkg/cli"
)

func main() {
	courseAndProofNum := strings.ToLower(cli.Prompt("Enter the target course (CC or ES) with number (1, 2, 3)"))
	course := fmt.Sprintf("%s.txt", courseAndProofNum[:2])

	if course != "cc.txt" && course != "es.txt" {
		log.Fatalf("Invalid course: %s. Please enter 'CC' or 'ES'.\n", courseAndProofNum)
		return
	}
	contentFromCourse, err := files.ReadFile(course)

	if err != nil {
		log.Fatalln("Error reading file:", err)
		return
	}

	for _, record := range contentFromCourse {
		path := fmt.Sprintf("internal/files/%s/", courseAndProofNum)
		_, err := files.GetAbsolutePath(path, record[1])

		if err != nil {
			log.Fatalln("Error reading file:", err)
			return
		}
	}
	log.Println("All files have been processed!")
}
