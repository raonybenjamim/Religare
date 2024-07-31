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

	generatorType, validationBypass, err := helpers.GetExecutionConfig()

	if err != nil {
		panic(err)
	}

	var signalGenerator converter.SignalGenerator
	var channelInterpreter interpreter.ChannelInterpreter

	switch generatorType {
	case models.RandomGeneratorType:
		signalGenerator = converter.NewRandomSignalGenerator(models.ConverterChannelSize)

	case models.WifiGeneratorType:
		signalGenerator = converter.NewWifiSignalGenerator(models.ConverterChannelSize, models.WifiThreshold)

	case models.TextInputGeneratorType:
		signalGenerator = converter.NewTextInputSignalGenerator(models.ConverterChannelSize)

	default:
		panic(fmt.Sprintf("No generator type selected. Value was: %v", generatorType.String()))
	}

	// Generate Signal
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
