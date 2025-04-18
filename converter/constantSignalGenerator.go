/*
 * Religare - An Instrumental Trans Communication solution for communicating with paranormal entities
 *
 * Copyright (C) 2024 Raony Benjamim
 * Check the LICENSE.md file for more information regarding the code license
 */
package converter

import "religare/models"

type ConstantSignalGenerator struct {
	channel    chan models.Binary
	BufferSize int
}

func NewConstantSignalGenerator(bufferSize int) *ConstantSignalGenerator {
	return &ConstantSignalGenerator{
		channel:    make(chan models.Binary, bufferSize),
		BufferSize: bufferSize,
	}
}

func (csg *ConstantSignalGenerator) GenerateSignal() {
	for {
		csg.channel <- models.One
	}
}

func (rsg *ConstantSignalGenerator) GetChannel() chan models.Binary {
	return rsg.channel
}

func (rsg *ConstantSignalGenerator) StopSignalGeneration() {
	close(rsg.channel)
}
