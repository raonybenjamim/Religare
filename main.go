package main

import (
	"flag"
	"fmt"
	"lazarus/converter"
	"lazarus/interpreter"
	"lazarus/models"
)

func main() {

	generatorType := flag.String(
		"generator-type",
		models.GeneratorType.Wifi,
		fmt.Sprintf("Define what will be the data generation behavior. Valid values are: %v", models.GeneratorType))

	// THIS WAY WE READ SIGNALS FROM THE WIFI
	// binaryChannel := make(chan models.Binary, models.ConverterChannelSize)

	// go signalConverter.AddBinaryData(0.75, binaryChannel)

	flag.Parse()

	var signalGenerator converter.SignalGenerator

	switch generatorType {
	case &models.GeneratorType.Random:
		signalGenerator = converter.NewRandomSignalGenerator(200)

	}

	// Generate Signal
	go signalGenerator.GenerateSignal()

	channelReader := interpreter.ChannelReader{
		Channel: signalGenerator.GetChannel(),
	}

	channelReader.ReadChannel()
}
