package main

import (
	"lazarus/converter"
	"lazarus/interpreter"
)

func main() {

	// THIS WAY WE READ SIGNALS FROM THE WIFI
	// binaryChannel := make(chan models.Binary, models.ConverterChannelSize)

	// go signalConverter.AddBinaryData(0.75, binaryChannel)

	// READING RANDOM DATA
	randomSignalGenerator := converter.NewRandomSignalGenerator(200)

	go randomSignalGenerator.GenerateSignal()

	interpreter.ReadChannel(randomSignalGenerator.Channel)
}
