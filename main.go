package main

import (
	signalConverter "lazarus/converter"
	"lazarus/interpreter"
	"lazarus/models"
)

func main() {

	binaryChannel := make(chan models.Binary, models.ConverterChannelSize)

	go signalConverter.AddBinaryData(0.75, binaryChannel)

	interpreter.ReadChannel(binaryChannel)
}
