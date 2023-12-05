package helpers

import (
	"log"
	"os"
	"strings"
)

func Read(filename string) []string {
	contents, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	stringContent := strings.Trim(string(contents), "\n")

	return strings.Split(stringContent, "\n")
}
