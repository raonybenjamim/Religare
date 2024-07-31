/*
 * Religare - An Instrumental Trans Communication solution for communicating with paranormal entities
 *
 * Copyright (C) 2024 Raony Benjamim
 * Check the LICENSE.md file for more information regarding the code license
 */
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
