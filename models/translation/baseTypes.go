package translation

import (
	"religare/config"
	"religare/models"
)

type TextIndex int

type Texts map[TextIndex]string

type LanguageMap map[models.Language]Texts

func (l LanguageMap) GetValue(index TextIndex) string {
	return l[config.AppLanguage][index]
}
