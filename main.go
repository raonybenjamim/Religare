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
	"flag"
	"fmt"
	"religare/config"
	"religare/converter"
	"religare/helpers"
	"religare/interpreter"
	"religare/models"
)

func main() {

	helpers.PrintLicense()
	config.AppLanguage = helpers.ChoseLanguage()

	helpers.PrintInitialMessage()

	generatorType := flag.String(
		"generator-type",
		models.GeneratorType.Wifi,
		fmt.Sprintf("Define what will be the data generation behavior. Valid values are: %v", models.GeneratorType))

	noValidation := flag.Bool(
		"no-validation",
		false,
		"If 'no-valiation' is true, the application will simply show any data received on the screen without any validation")

	flag.Parse()

	var signalGenerator converter.SignalGenerator
	var channelInterpreter interpreter.ChannelInterpreter

	switch *generatorType {
	case models.GeneratorType.Random:
		signalGenerator = converter.NewRandomSignalGenerator(models.ConverterChannelSize)

	case models.GeneratorType.Wifi:
		signalGenerator = converter.NewWifiSignalGenerator(models.ConverterChannelSize, models.WifiThreshold)

	case models.GeneratorType.TextInput:
		signalGenerator = converter.NewTextInputSignalGenerator(models.ConverterChannelSize)

	default:
		panic(fmt.Sprintf("No generator type selected. Value was: %v", *generatorType))
	}

	// Generate Signal
	go signalGenerator.GenerateSignal()

	if *noValidation {
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
