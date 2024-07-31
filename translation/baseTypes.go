package translation

import (
	"religare/config"
	"religare/customTypes"
)

type TextIndex int

type Texts map[TextIndex]string

type LanguageMap map[customTypes.Language]Texts

func (l LanguageMap) GetValue(index TextIndex) string {
	return l[config.AppLanguage][index]
}
