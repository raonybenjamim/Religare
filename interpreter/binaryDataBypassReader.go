/*
 * Religare - An Instrumental Trans Communication solution for communicating with paranormal entities
 *
 * Copyright (C) 2024 Raony Benjamim
 * Check the LICENSE.md file for more information regarding the code license
 */
package interpreter

import "lazarus/models"

type BinaryDataBypassReader struct {
	Channel <-chan models.Binary
}

func (reader *BinaryDataBypassReader) GetChannel() <-chan models.Binary {
	return reader.Channel
}

func (reader *BinaryDataBypassReader) ReadChannel() {

}
