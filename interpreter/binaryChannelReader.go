/*
 * Religare - An Instrumental Trans Communication solution for communicating with paranormal entities
 *
 * Copyright (C) 2024 Raony Benjamim
 * Check the LICENSE.md file for more information regarding the code license
 */
package interpreter

import (
	"errors"
	"fmt"
	"lazarus/helpers"
	"lazarus/models"
	"strconv"
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

func (reader *BinaryDataInterpreter) binaryToText(binaryString string) (string, error) {
	var text string

	if len(binaryString)%8 != 0 {
		return "", errors.New("binary string length is not a multiple of 8")
	}

	for i := 0; i < len(binaryString); i += models.ByteSize {
		// Get the next 8 bits
		byteString := binaryString[i : i+models.ByteSize]

		// Convert the 8-bit binary string to a decimal (base 10) integer
		charCode, err := strconv.ParseInt(byteString, 2, 64)
		if err != nil {
			return "", err
		}

		// Convert the integer to a corresponding ASCII character
		text += string(rune(charCode))
	}

	return text, nil
}

func (reader *BinaryDataInterpreter) binaryToHex(binaryString string) (string, error) {
	if len(binaryString)%4 != 0 {
		return "", errors.New("binary string length is not a multiple of 4")
	}

	var hexString string

	for i := 0; i < len(binaryString); i += 4 {
		// Get the next 4 bits
		bitChunk := binaryString[i : i+4]

		// Convert the 4-bit binary string to a decimal (base 10) integer
		decimalValue, err := strconv.ParseInt(bitChunk, 2, 64)
		if err != nil {
			return "", err
		}

		// Convert the integer to a corresponding hexadecimal character
		hexString += strconv.FormatInt(decimalValue, 16)
	}

	return hexString, nil
}

func (reader *BinaryDataInterpreter) isValidStringMessage(expectedBinaryChecksum string, messageData string) bool {
	expectedHash, err := reader.binaryToHex(expectedBinaryChecksum)

	if err != nil {
		return false
	}

	messageHash := helpers.GetMd5HashFromString(messageData)

	return expectedHash == messageHash

}

func (reader *BinaryDataInterpreter) readTextMessageFromChannel(headers models.MessageHeaders) (string, error) {
	textContent, err := reader.binaryToText(helpers.GetDataFromChannel(reader.Channel, headers.MessageSizeBytes*models.ByteSize))

	if err != nil {
		return "", fmt.Errorf("Error while reading binary data form channel: " + err.Error())
	}

	if !reader.isValidStringMessage(headers.Checksum, textContent) {
		return "", fmt.Errorf("the received message was not valid")
	}

	return textContent, nil
}
