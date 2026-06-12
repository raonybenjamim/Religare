/*
 * Religare - An Instrumental Trans Communication solution for communicating with paranormal entities
 *
 * Copyright (C) 2024 Raony Benjamim
 * Check the LICENSE.md file for more information regarding the code license
 */
package reporting

import (
	"fmt"
	"os"
	"time"
)

// LogPositiveResult logs a successful message reception or event to a local file.
func LogPositiveResult(result string, mode string) {
	logMessage := fmt.Sprintf("[%s] [Mode: %s] Result: %s\n", time.Now().Format(time.RFC3339), mode, result)
	
	// Open the file in append mode. If it doesn't exist, create it with 0644 permissions.
	file, err := os.OpenFile("results_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open results_log.txt: %v\n", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(logMessage); err != nil {
		fmt.Printf("Failed to write to results_log.txt: %v\n", err)
	}
}
