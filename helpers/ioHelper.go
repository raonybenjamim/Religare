package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
)

func WriteStringToFile(content string) (string, error) {
	// Generate a new GUID
	fileName := uuid.New().String() + ".txt"

	// Write the content to the file
	err := os.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		return "", err
	}

	return fileName, nil
}

func HealthcheckPrint(message string) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Println(message)
	}
}
