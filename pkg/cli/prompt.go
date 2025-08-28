package cli

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Prompt(message string) string {
	fmt.Printf("%s: ", message)
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalf("Error reading input: %v", err)
		return ""
	}
	return strings.TrimSpace(input)
}
