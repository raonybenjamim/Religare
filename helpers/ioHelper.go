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
	"religare/customTypes"
	"religare/models"
	"religare/translation"
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

func ChoseLanguage() customTypes.Language {
	var response string

	fmt.Println("0 for Portuguese 1 for English")
	fmt.Scan(&response)

	language, err := strconv.Atoi(response)

	if err != nil {
		panic(err)
	}

	return customTypes.Language(language)
}

func GetDataFromChannel(channel <-chan models.Binary, quantity int) string {
	var builder strings.Builder

	for i := 0; i < quantity; i++ {
		binaryValue, ok := <-channel

		if !ok {
			panic(translation.ErrorTexts.GetValue(translation.ChannelClosedError))
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

func PrintInitialMessage() {
	fmt.Println(translation.MenuTexts.GetValue(translation.WelcomeMessageMenu))
}

func GetExecutionConfig() (models.GeneratorType, bool, error) {
	var chosenGenerator string
	var chosenBypass string

	fmt.Println(translation.MenuTexts.GetValue(translation.SelectGeneratorMenu))
	fmt.Scan(&chosenGenerator)

	fmt.Println(translation.MenuTexts.GetValue(translation.BypassValidationMenu))
	fmt.Scan(&chosenBypass)

	generatorType, err := strconv.Atoi(chosenGenerator)

	if err != nil {
		return 0, false, err
	}

	shouldBypass := parseByass(chosenBypass)

	return models.GeneratorType(generatorType), shouldBypass, nil

}

func parseByass(chosenBypass string) bool {
	switch chosenBypass {
	case "1":
		return true
	default:
		return false
	}
}
