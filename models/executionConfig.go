/*
 * Religare - An Instrumental Trans Communication solution for communicating with paranormal entities
 *
 * Copyright (C) 2024 Raony Benjamim
 * Check the LICENSE.md file for more information regarding the code license
 */
package models

import (
	"religare/customTypes"
	"strings"
)

type ExecutionConfig struct {
	Language               string                              `json:"language"`
	GeneratorType          string                              `json:"generatorType"`
	ValidationBypass       bool                                `json:"validationBypass"`
	WebSocketConfig        customTypes.WebsocketConnectionInfo `json:"webSocketConfig"`
	ScreenExhibitionConfig customTypes.ScreenExhibitionConfig  `json:"screenExhibitionConfig"`
}

func (e *ExecutionConfig) ParseLanguage() customTypes.Language {
	switch strings.ToLower(e.Language) {
	case "pt-br":
		return customTypes.Portuguese
	case "en":
		return customTypes.English
	default:
		return customTypes.English
	}
}

func (e *ExecutionConfig) ParseGeneratorType() customTypes.GeneratorType {
	switch strings.ToLower(e.GeneratorType) {
	case "wifisignal":
		return customTypes.WifiGeneratorType
	case "random":
		return customTypes.RandomGeneratorType
	case "textinput":
		return customTypes.TextInputGeneratorType
	case "datasender":
		return customTypes.DataSenderGeneratorType
	case "datareceiver":
		return customTypes.DataReceiverGeneratorType
	default:
		return customTypes.WifiGeneratorType
	}
}
