package converter

import (
	"lazarus/models"
	"math/rand"
)

type RandomSignalGenerator struct {
	Channel    chan models.Binary
	BufferSize int
}

func NewRandomSignalGenerator(bufferSize int) *RandomSignalGenerator {
	return &RandomSignalGenerator{
		Channel:    make(chan models.Binary, bufferSize),
		BufferSize: bufferSize,
	}
}

func (RandomSignalGenerator) generateBinary() models.Binary {
	randomNumber := rand.Int()

	return models.Binary(randomNumber % 2)
}

func (rsg *RandomSignalGenerator) GenerateSignal() {
	for {
		rsg.Channel <- rsg.generateBinary()
	}
}

func (rsg *RandomSignalGenerator) GetChannel() chan models.Binary {
	return rsg.Channel
}
