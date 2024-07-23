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
	"os/signal"
	"religare/models"
	"strconv"
	"strings"
	"syscall"
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
	fmt.Println("Copyright (C) 2024 - Raony Benjamim")
	fmt.Println("This program comes with ABSOLUTELY NO WARRANTY;")
	fmt.Println("This is free software, and you are welcome to redistribute it")
	fmt.Println("under certain conditions; Check https://www.gnu.org/licenses/ for details.")
}

func ChoseLanguage() models.Language {
	var response string

	fmt.Println("0 for Portuguese 1 for English")
	fmt.Scan(response)

	language, err := strconv.Atoi(response)

	if err != nil {
		panic(err)
	}

	return models.Language(language)
}

func GetDataFromChannel(channel <-chan models.Binary, quantity int) string {
	var builder strings.Builder

	for i := 0; i < quantity; i++ {
		binaryValue, ok := <-channel

		if !ok {
			panic("Channel closed before enough data was provided")
		}

		switch binaryValue {
		case models.One:
			builder.WriteString("1")
		case models.Zero:
			builder.WriteString("0")
		default:
			panic("Channel data is neither 0 or 1")
		}
	}

	return builder.String()
}

func SetupKeyboardSignalHandling() {
	keyboardSignalChannel := make(chan os.Signal, 1)

	signal.Notify(keyboardSignalChannel, os.Interrupt, syscall.SIGTERM)

	go HandleKeyboardSignals(keyboardSignalChannel)

}

func HandleKeyboardSignals(keyboardSignalChannel chan os.Signal) {
	for {
		receivedSignal := <-keyboardSignalChannel

		switch receivedSignal {
		case os.Interrupt:
			// handle the Ctrl+C signal
		case syscall.SIGTERM:
			os.Exit(0)
		}
	}

}
