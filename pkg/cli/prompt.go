package cli

import (
	"fmt"
	"log"
)

func Prompt(message string) string {
	var input string
	fmt.Printf("%s: ", message)
	_, err := fmt.Scanln(&input)

	if err != nil {
		log.Fatalf("Error reading input: %v", err)
		return ""
	}
	return input
}
