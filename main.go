package main

import (
	"fmt"
	signalInterpreter "lazarus/interpreter"
	"lazarus/models"
)

func main() {

	binaryChannel := make(chan models.Binary, models.InterpreterChannelSize)

	go signalInterpreter.AddBinaryData(0.5, binaryChannel)

	for binary := range binaryChannel {
		fmt.Printf("Binary signal strength: %s\n", binary)
	}

	fmt.Println("Channel closed")
}
