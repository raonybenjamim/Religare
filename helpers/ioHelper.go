/*
 * Religare - An Instrumental Trans Communication solution for communicating with paranormal entities
 *
 * Copyright (C) 2024 Raony Benjamim
 * Check the LICENSE.md file for more information regarding the code license
 */
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

func PrintLicense() {
	fmt.Println("Copyright (C) <year> <author name>")
	fmt.Println("This program comes with ABSOLUTELY NO WARRANTY;")
	fmt.Println("This is free software, and you are welcome to redistribute it")
	fmt.Println("under certain conditions; Check https://www.gnu.org/licenses/ for details.")
}
