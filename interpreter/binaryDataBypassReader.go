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
	"time"
)

type BinaryDataBypassReader struct {
	Channel <-chan models.Binary
}

func (reader *BinaryDataBypassReader) GetChannel() <-chan models.Binary {
	return reader.Channel
}

func (reader *BinaryDataBypassReader) ReadChannel() {
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		binaryCharacter := helpers.GetDataFromChannel(reader.Channel, 4*models.ByteSize)

		textCharacter, err := helpers.BinaryStringToString(binaryCharacter)

		if err != nil {
			fmt.Println("error while reading binary data form channel: " + err.Error())
		}
		fmt.Print(textCharacter)
	}
}
