package converter

import "lazarus/models"

type SignalGenerator interface {
	GetChannel() chan models.Binary

	GenerateSignal()
}
