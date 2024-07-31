/*
 * Religare - An Instrumental Trans Communication solution for communicating with paranormal entities
 *
 * Copyright (C) 2024 Raony Benjamim
 * Check the LICENSE.md file for more information regarding the code license
 */
package translation

import "religare/customTypes"

const (
	ChannelClosedError TextIndex = iota
	WrongInputError
)

var ErrorTexts = LanguageMap{
	customTypes.English: Texts{
		ChannelClosedError: "Channel closed before enough data was provided",
		WrongInputError:    "Wrong input received: ",
	},
	customTypes.Portuguese: Texts{
		ChannelClosedError: "Canal foi fechado antes de dados suficientes terem sido providos",
		WrongInputError:    "Entrada inválida: ",
	},
}
