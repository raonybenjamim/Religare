/*
 * Religare - An Instrumental Trans Communication solution for communicating with paranormal entities
 *
 * Copyright (C) 2024 Raony Benjamim
 * Check the LICENSE.md file for more information regarding the code license
 */
package interpreter

import (
	"fmt"
	"religare/helpers"
	"religare/models"
)

type BinaryDataInterpreter struct {
	Channel <-chan models.Binary
}

func (reader *BinaryDataInterpreter) GetChannel() <-chan models.Binary {
	return reader.Channel
}
func (reader *BinaryDataInterpreter) ReadChannel() {

	// go helpers.HealthcheckPrint("Reading messages")

	for {
		if !(helpers.GetDataFromChannel(reader.Channel, models.ValidStartBits) == models.ValidStart) {
			continue
		}

		messaageHeaders, err := reader.readHeadersFromChannel()

		if err != nil {
			continue
		}

		// Read the "size" bits from the message and convert it to
		// a text, image or audio

		switch messaageHeaders.MessageType {
		case models.MessageType.Text:
			textContent, err := reader.readTextMessageFromChannel(messaageHeaders)

			if err == nil {
				println("Got message: " + textContent)
			}

		case models.MessageType.TextFile:
			textContent, err := reader.readTextMessageFromChannel(messaageHeaders)

			if err == nil {
				filename, err := helpers.WriteStringToFile(textContent)

				if err != nil {
					break
				}

				println("File Received: ", filename)
			}

		}
	}
}

func (reader *BinaryDataInterpreter) readHeadersFromChannel() (models.MessageHeaders, error) {
	messageType := helpers.GetDataFromChannel(reader.Channel, models.MessageTypeBits)

	// Only accept valid message types
	if !helpers.IsValidMessageHeader(messageType) {
		return models.MessageHeaders{}, fmt.Errorf("wrong message type received: %v", messageType)
	}

	checksumData := helpers.GetDataFromChannel(reader.Channel, models.ChecksumBits)

	messageSize, err := helpers.ConvertBinaryToInt(helpers.GetDataFromChannel(reader.Channel, models.MessageSizeBits))

	if err != nil {
		return models.MessageHeaders{}, err
	}

	return models.MessageHeaders{
		MessageType:      messageType,
		Checksum:         checksumData,
		MessageSizeBytes: messageSize,
	}, nil
}

func (reader *BinaryDataInterpreter) isValidStringMessage(expectedBinaryChecksum string, messageData string) bool {
	expectedHash, err := helpers.BinaryStringToHexString(expectedBinaryChecksum)

	if err != nil {
		return false
	}

	messageHash := helpers.GetMd5HashFromString(messageData)

	return expectedHash == messageHash

}

func (reader *BinaryDataInterpreter) readTextMessageFromChannel(headers models.MessageHeaders) (string, error) {
	textContent, err := helpers.BinaryStringToString(helpers.GetDataFromChannel(reader.Channel, headers.MessageSizeBytes*models.ByteSize))

	if err != nil {
		return "", fmt.Errorf("Error while reading binary data form channel: " + err.Error())
	}

	if !reader.isValidStringMessage(headers.Checksum, textContent) {
		return "", fmt.Errorf("the received message was not valid")
	}

	return textContent, nil
}
