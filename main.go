package main

import (
	"fmt"
	signalConverter "lazarus/converter"
	"lazarus/models"
)

func main() {

	binaryChannel := make(chan models.Binary, models.ConverterChannelSize)

	go signalConverter.AddBinaryData(0.5, binaryChannel)

	for binary := range binaryChannel {
		fmt.Printf("Binary signal strength: %s\n", binary)
	}

	fmt.Println("Channel closed")
}
