/*
 * Religare - An Instrumental Trans Communication solution for communicating with paranormal entities
 *
 * Copyright (C) 2024 Raony Benjamim
 * Check the LICENSE.md file for more information regarding the code license
 */
package helpers

import (
	"crypto/md5"
	"fmt"
	"lazarus/models"
)

func IsValidMessageHeader(header string) bool {
	switch header {
	case models.MessageType.Text,
		models.MessageType.TextFile,
		models.MessageType.Image,
		models.MessageType.ImageFile,
		models.MessageType.Audio,
		models.MessageType.AudioFile:
		return true
	default:
		return false
	}
}

func GetMd5HashFromString(content string) string {
	data := []byte(content)

	return fmt.Sprintf("%x", md5.Sum(data))
}

func GetMd5HashFromBytes(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}
