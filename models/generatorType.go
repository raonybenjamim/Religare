/*
 * Religare - An Instrumental Trans Communication solution for communicating with paranormal entities
 *
 * Copyright (C) 2024 Raony Benjamim
 * Check the LICENSE.md file for more information regarding the code license
 */
package models

import (
	"errors"
	"religare/customTypes"
	"religare/translation"
)

func ValidateGeneratorType(g customTypes.GeneratorType) error {
	if g > customTypes.TextInputGeneratorType {
		return errors.New(translation.ErrorTexts.GetValue(translation.WrongInputError) + g.String())
	}
	return nil
}
