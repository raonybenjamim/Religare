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

	flag.Parse()

	var signalGenerator converter.SignalGenerator

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

	channelReader := interpreter.ChannelReader{
		Channel: signalGenerator.GetChannel(),
	}

	channelReader.ReadChannel()
}
