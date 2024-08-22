/*
 * Religare - An Instrumental Trans Communication solution for communicating with paranormal entities
 *
 * Copyright (C) 2024 Raony Benjamim
 * Check the LICENSE.md file for more information regarding the code license
 */
package converter

import (
	"os"
	"religare/helpers"
	"religare/models"
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
		fileContent, err := os.ReadFile("./models/sampleTextContent")
		if err != nil {
			panic("fatal failure while reading the sample file: " + err.Error())
		}

		message := string(fileContent)

		checkSum := helpers.GetMd5HashFromString(message)

		binaryChecksum, err := helpers.HexTo4BitBinary(checkSum)

		if err != nil {
			panic("fatal failure while generating message md5")
		}

		messageBinary := helpers.StringToBinaryString(message)

		// create valid headers
		headers := models.ValidStart +
			models.MessageType.Text +
			binaryChecksum +
			helpers.IntToBinaryString(len(messageBinary)/8, 10)

		binaryData := helpers.BinaryStringToBinaryData(headers + messageBinary)

		for _, bit := range binaryData {
			tsg.channel <- bit
		}
	}
}

func (tsg *TextInputSignalGenerator) StopSignalGeneration() {
	close(tsg.channel)
}

func (tsg *TextInputSignalGenerator) GetChannel() chan models.Binary {
	return tsg.channel
}
