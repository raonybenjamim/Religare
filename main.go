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
	"religare/converter"
	"religare/customTypes"
	"religare/interpreter"
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

	var executionConfig models.ExecutionConfig
	var generatorType customTypes.GeneratorType
	var validationBypass bool
	var err error

	configContent, err := os.ReadFile(configFilePath)

	if err != nil {
		panic(fmt.Sprintf("%s not found", configFilePath))
	}

	err = json.Unmarshal(configContent, &executionConfig)

	if err != nil {
		panic(err)
	}

	config.AppLanguage = executionConfig.ParseLanguage()
	generatorType = executionConfig.ParseGeneratorType()
	validationBypass = executionConfig.ValidationBypass
	config.WebSocketConfig = &executionConfig.WebSocketConfig

	var signalGenerator converter.SignalGenerator
	var channelInterpreter interpreter.ChannelInterpreter

	switch generatorType {
	case customTypes.RandomGeneratorType:
		signalGenerator = converter.NewRandomSignalGenerator(models.ConverterChannelSize)

	case customTypes.WifiGeneratorType:
		signalGenerator = converter.NewWifiSignalGenerator(models.ConverterChannelSize, models.WifiThreshold)

	case customTypes.TextInputGeneratorType:
		signalGenerator = converter.NewTextInputSignalGenerator(models.ConverterChannelSize)

	case customTypes.DataSenderGeneratorType:
		signalGenerator = converter.NewDataSenderGenerator(models.ConverterChannelSize)

	case customTypes.DataReceiverGeneratorType:
		signalGenerator = converter.NewDataReceiverGenerator(models.ConverterChannelSize)

	default:
		panic(fmt.Sprintf("No generator type selected. Value was: %v", generatorType.String()))
	}

	// Generate Signal
	if generatorType == customTypes.DataSenderGeneratorType {
		signalGenerator.GenerateSignal()
	}

	go signalGenerator.GenerateSignal()

	if validationBypass {
		channelInterpreter = &interpreter.BinaryDataBypassReader{
			Channel: signalGenerator.GetChannel(),
		}
	} else {
		channelInterpreter = &interpreter.BinaryDataInterpreter{
			Channel: signalGenerator.GetChannel(),
		}
	}

	channelInterpreter.ReadChannel()
}
