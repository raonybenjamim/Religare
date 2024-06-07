package helpers

import "lazarus/models"

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
