/*
 * Religare - An Instrumental Trans Communication solution for communicating with paranormal entities
 *
 * Copyright (C) 2024 Raony Benjamim
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program. If not, see <https://www.gnu.org/licenses/>.
 */
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"religare/config"
	"religare/models"
)

var configFilePath string

func init() {
	flag.StringVar(
		&configFilePath,
		"config",
		"./runConfigurations/randomNoValidation.json",
		"provide a custom config file for starting the application")
}

func main() {

	flag.Parse()

	parseConfig(configFilePath)

	switch config.RunMode {
	case config.SenderMode:
		// Execute Send Data Function
		sendData()
	case config.ReceiverMode:
		// Execute Receive Data Function
		receiveData()
	case config.ExecuteLocalMode:
		executeLocal()
	}
}

func parseConfig(configFilePath string) {

	var executionConfig models.ExecutionConfig

	configContent, err := os.ReadFile(configFilePath)

	if err != nil {
		panic(fmt.Sprintf("%s not found", configFilePath))
	}

	err = json.Unmarshal(configContent, &executionConfig)

	if err != nil {
		panic(fmt.Sprintf("Error while parsing config: %v", err))
	}

	config.AppLanguage = executionConfig.ParseLanguage()
	config.Calibration = executionConfig.CalibrationConfig
	config.WebSocketConnection = executionConfig.WebSocketConnectionConfig
	config.RunMode = executionConfig.RunMode
}

func sendData() {
	// Load the data producer
	// Check if it is a constant data sender
	// Start the producer
}

func receiveData() {
	// Load the data receiver
	// Check the type of bypass
}

func executeLocal() {
	// This mode both produce and consumes data with local channels
}
