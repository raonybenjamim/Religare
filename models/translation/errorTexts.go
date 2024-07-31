package translation

import "religare/models"

const (
	ChannelClosedError TextIndex = iota
	WrongInputError
)

var ErrorTexts = LanguageMap{
	models.English: Texts{
		ChannelClosedError: "Channel closed before enough data was provided",
		WrongInputError:    "Wrong input received: ",
	},
	models.Portuguese: Texts{
		ChannelClosedError: "Canal foi fechado antes de dados suficientes terem sido providos",
		WrongInputError:    "Entrada inválida: ",
	},
}
