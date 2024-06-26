/*
 * Religare - An Instrumental Trans Communication solution for communicating with paranormal entities
 *
 * Copyright (C) 2024 Raony Benjamim
 * Check the LICENSE.md file for more information regarding the code license
 */

package interpreter

import "lazarus/models"

type ChannelInterpreter interface {
	ReadChannel()
	GetChannel() chan models.Binary
}
