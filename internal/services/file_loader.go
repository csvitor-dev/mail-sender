package services

import (
	"github.com/csvitor-dev/mail-sender/internal/entities"
)

func LoadFiles(jobs chan<- *entities.EmailJob) error {
	contentFromCC, err := readFile("cc.txt")

	if err != nil {
		return err
	}
	contentFromES, err := readFile("es.txt")

	if err != nil {
		return err
	}
	bodyMessage := "Segue em anexo a terceira prova de Matemática Discreta.\n\nAtenciosamente,\nVitor Costa de Sousa"

	for _, record := range contentFromCC {
		fileToAttach, err := GetAbsolutePath("internal/files/cc3/", record[1])

		if err != nil {
			return err
		}
		jobs <- entities.NewEmail(record[0], "Prova de Matemática Discreta [CC 2025.1]", bodyMessage, "text/plain", fileToAttach)
	}

	for _, record := range contentFromES {
		fileToAttach, err := GetAbsolutePath("internal/files/es3/", record[1])

		if err != nil {
			return err
		}
		jobs <- entities.NewEmail(record[0], "Prova de Matemática Discreta [ES 2025.1]", bodyMessage, "text/plain", fileToAttach)
	}
	return nil
}
