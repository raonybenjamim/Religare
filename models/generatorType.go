package models

import (
	"errors"
	"religare/models/translation"
)

type GeneratorType int

const (
	RandomGeneratorType GeneratorType = iota
	WifiGeneratorType
	TextInputGeneratorType
)

func (g GeneratorType) String() string {
	switch g {
	case RandomGeneratorType:
		return "random"
	case WifiGeneratorType:
		return "wifi"
	case TextInputGeneratorType:
		return "textInput"
	default:
		return "wifi"
	}
}

func (g GeneratorType) Validate() error {
	if g > TextInputGeneratorType {
		return errors.New(translation.ErrorTexts.GetValue(translation.WrongInputError) + string(g))
	}
	return nil
}
