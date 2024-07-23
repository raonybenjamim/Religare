package translation

import "religare/models"

const (
	ChannelClosedError TextIndex = iota
)

var ErrorTexts = LanguageMap{
	models.English: Texts{
		ChannelClosedError: "Channel closed before enough data was provided",
	},
}
