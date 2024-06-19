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

		binaryChecksum, err := helpers.HexTo4BitBinary(checkSum)

		if err != nil {
			panic("fatal failure while generating message md5")
		}

		// create valid headers
		headers := models.ValidStart +
			models.MessageType.Text +
			binaryChecksum +
			helpers.IntToBinaryString(len(message), 10)

		messageBinary := helpers.StringToBinaryString(message, 8)

		binaryData := helpers.BinaryStringToBinaryData(headers + messageBinary)

		for _, bit := range binaryData {
			tsg.channel <- bit
		}
	}
}

func (tsg *TextInputSignalGenerator) GetChannel() chan models.Binary {
	return tsg.channel
}
