package interpreter

import (
	"errors"
	"fmt"
	"lazarus/helpers"
	"lazarus/models"
	"strconv"
	"strings"
)

type ChannelReader struct {
	Channel <-chan models.Binary
}

func (reader *ChannelReader) ReadChannel() {

	go helpers.HealthcheckPrint("Reading messages")

	for {
		if !(reader.getDataFromChannel(models.ValidStartBits) == models.ValidStart) {
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

func (reader *ChannelReader) readHeadersFromChannel() (models.MessageHeaders, error) {
	messageType := reader.getDataFromChannel(models.MessageTypeBits)

	// Only accept valid message types
	if !helpers.IsValidMessageHeader(messageType) {
		return models.MessageHeaders{}, fmt.Errorf("wrong message type received: %v", messageType)
	}

	checksumData := reader.getDataFromChannel(models.ChecksumBits)

	messageSize, err := helpers.ConvertBinaryToInt(reader.getDataFromChannel(models.MessageSizeBits))

	if err != nil {
		return models.MessageHeaders{}, err
	}

	return models.MessageHeaders{
		MessageType:      messageType,
		Checksum:         checksumData,
		MessageSizeBytes: messageSize,
	}, nil
}

func (reader *ChannelReader) getDataFromChannel(quantity int) string {
	var builder strings.Builder

	for i := 0; i < quantity; i++ {
		binaryValue, ok := <-reader.Channel

		if !ok {
			panic("Channel closed before enough data was provided")
		}

		switch binaryValue {
		case models.One:
			builder.WriteString("1")
		case models.Zero:
			builder.WriteString("0")
		default:
			panic("Channel data is neither 0 or 1")
		}
	}

	return builder.String()
}

func (reader *ChannelReader) binaryToText(binaryString string) (string, error) {
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

func (reader *ChannelReader) isValidStringMessage(expectedBinaryChecksum string, messageData string) bool {
	expectedHash, err := reader.binaryToText(expectedBinaryChecksum)

	if err != nil {
		return false
	}

	messageHash := helpers.GetMd5HashFromString(messageData)

	return expectedHash == messageHash

}

func (reader *ChannelReader) readTextMessageFromChannel(headers models.MessageHeaders) (string, error) {
	textContent, err := reader.binaryToText(reader.getDataFromChannel(headers.MessageSizeBytes * models.ByteSize))

	if err != nil {
		return "", fmt.Errorf("Error while reading binary data form channel: " + err.Error())
	}

	if !reader.isValidStringMessage(headers.Checksum, textContent) {
		return "", fmt.Errorf("the received message was not valid")
	}

	return textContent, nil
}
