package main

import (
	"fmt"
	signalInterpreter "lazarus/interpreter"
	"lazarus/models"
)

func main() {

	ch := make(chan models.Binary, models.InterpreterChannelSize)

	go signalInterpreter.AddBinaryData(0.5, ch)

	for binary := range ch {
		fmt.Printf("Binary signal strength: %s\n", binary)
	}

	fmt.Println("Channel closed")
}
