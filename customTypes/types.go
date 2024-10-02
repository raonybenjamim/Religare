/*
 * Religare - An Instrumental Trans Communication solution for communicating with paranormal entities
 *
 * Copyright (C) 2024 Raony Benjamim
 * Check the LICENSE.md file for more information regarding the code license
 */
package customTypes

type Language int

const (
	Portuguese Language = iota
	English
)

func (l Language) String() string {
	switch l {
	case Portuguese:
		return "Portuguese"

	case English:
		return "English"

	default:
		return "unknown"
	}
}

type GeneratorType int

const (
	WifiGeneratorType GeneratorType = iota
	RandomGeneratorType
	TextInputGeneratorType
	DataSenderGeneratorType
	DataReceiverGeneratorType
)

func (g GeneratorType) String() string {
	switch g {
	case RandomGeneratorType:
		return "random"
	case WifiGeneratorType:
		return "wifi"
	case TextInputGeneratorType:
		return "textInput"
	case DataSenderGeneratorType:
		return "dataSender"
	case DataReceiverGeneratorType:
		return "dataReceiver"
	default:
		return "wifi"
	}
}

type WebsocketConnectionInfo struct {
	Ip   string `json:"ip"`
	Port int    `json:"port"`
}
