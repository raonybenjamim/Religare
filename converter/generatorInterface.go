/*
 * Religare - An Instrumental Trans Communication solution for communicating with paranormal entities
 *
 * Copyright (C) 2024 Raony Benjamim
 * Check the LICENSE.md file for more information regarding the code license
 */
package converter

import "religare/models"

type SignalGenerator interface {
	GetChannel() chan models.Binary

	GenerateSignal()

	StopSignalGeneration()
}
