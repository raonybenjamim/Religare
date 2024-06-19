package converter

import (
	"lazarus/models"
	"math/rand"
)

type RandomSignalGenerator struct {
	channel    chan models.Binary
	BufferSize int
}

func NewRandomSignalGenerator(bufferSize int) *RandomSignalGenerator {
	return &RandomSignalGenerator{
		channel:    make(chan models.Binary, bufferSize),
		BufferSize: bufferSize,
	}
}

func (RandomSignalGenerator) generateBinary() models.Binary {
	randomNumber := rand.Int()

	return models.Binary(randomNumber % 2)
}

func (rsg *RandomSignalGenerator) GenerateSignal() {
	for {
		rsg.channel <- rsg.generateBinary()
	}
}

func (rsg *RandomSignalGenerator) GetChannel() chan models.Binary {
	return rsg.channel
}
