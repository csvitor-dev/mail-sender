package files

import (
	"fmt"

	"github.com/csvitor-dev/mail-sender/internal/entities"
)

// LoadFromFiles: Loads email jobs from a specified course file and prepares them for sending.
func LoadFromFiles(courseAndProofNum, subject, bodyMessage string, jobs chan<- *entities.EmailJob) error {
	course := fmt.Sprintf("%s.txt", courseAndProofNum[:2])
	contentFromCourse, err := ReadFile(course)

	if err != nil {
		return err
	}

	for _, record := range contentFromCourse {
		path := fmt.Sprintf("internal/files/%s/", courseAndProofNum)
		fileToAttach, err := GetAbsolutePath(path, record[1])

		if err != nil {
			return err
		}
		job, err := entities.NewEmail(record[0], subject, bodyMessage, "application/pdf", fileToAttach)

		if err != nil {
			return err
		}
		jobs <- job
	}
	return nil
}
