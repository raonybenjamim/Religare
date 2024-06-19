package converter

import (
	"fmt"
	"lazarus/helpers"
	"lazarus/models"
)

type TextInputSignalGenerator struct {
	channel    chan models.Binary
	bufferSize int
}

func NewTextInputSignalGenerator(bufferSize int) *TextInputSignalGenerator {
	return &TextInputSignalGenerator{
		channel:    make(chan models.Binary, bufferSize),
		bufferSize: bufferSize,
	}
}

func (tsg *TextInputSignalGenerator) GenerateSignal() {
	for {
		var message string

		fmt.Println("Enter a message")
		fmt.Scan(&message)

		checkSum := helpers.GetMd5HashFromString(message)

		// create valid headers
		headers := models.ValidStart +
			models.MessageType.Text +
			helpers.StringToBinaryString(checkSum, 4) +
			helpers.IntToBinaryString(len(message), 10)

		messageBinary := helpers.StringToBinaryString(message, 8)

		fmt.Println(headers + messageBinary)
		fmt.Println("Checksum size: ", len(helpers.StringToBinaryString(checkSum, 4)))
		fmt.Println("checksum: ", checkSum, "size: ", len(checkSum))

		for _, bit := range helpers.BinaryStringToBinaryData(headers + messageBinary) {
			tsg.channel <- bit
		}
	}
}

func (tsg *TextInputSignalGenerator) GetChannel() chan models.Binary {
	return tsg.channel
}
