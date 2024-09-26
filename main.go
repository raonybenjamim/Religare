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
	"religare/customTypes"
	"religare/helpers"
	"religare/interpreter"
	"religare/models"
)

var debug bool

func init() {
	flag.BoolVar(&debug, "debug", false, "enable debug mode")
}

func main() {

	flag.Parse()

	var generatorType customTypes.GeneratorType
	var validationBypass bool
	var err error

	if !debug {
		helpers.PrintLicense()
		config.AppLanguage = helpers.ChoseLanguage()

		helpers.PrintInitialMessage()

		generatorType, validationBypass, err = helpers.GetExecutionConfig()

		if err != nil {
			panic(err)
		}
	} else {
		config.AppLanguage = customTypes.English
		generatorType = customTypes.TextInputGeneratorType
		validationBypass = false
	}

	var signalGenerator converter.SignalGenerator
	var channelInterpreter interpreter.ChannelInterpreter

	switch generatorType {
	case customTypes.RandomGeneratorType:
		signalGenerator = converter.NewRandomSignalGenerator(models.ConverterChannelSize)

	case customTypes.WifiGeneratorType:
		signalGenerator = converter.NewWifiSignalGenerator(models.ConverterChannelSize, models.WifiThreshold)

	case customTypes.TextInputGeneratorType:
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
