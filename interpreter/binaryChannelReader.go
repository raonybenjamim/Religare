package interpreter

import (
	"fmt"
	"lazarus/helpers"
	"lazarus/models"
	"strings"
)

type ChannelReader struct {
	Channel <-chan models.Binary
}

func (reader *ChannelReader) ReadChannel() bool {
	for {
		if !(reader.getDataFromChannel(models.ValidStartBits) == models.ValidStart) {
			continue
		}

		messaageHeaders, err := reader.readHeadersFromChannel()

		if err != nil {
			continue
		}

		fmt.Printf("Headers found! %v", messaageHeaders)
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
