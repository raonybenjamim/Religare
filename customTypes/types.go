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
	DataReceiverGeneratorType
	ConstantSignalGeneratorType
)

func (g GeneratorType) String() string {
	switch g {
	case RandomGeneratorType:
		return "random"
	case WifiGeneratorType:
		return "wifi"
	case TextInputGeneratorType:
		return "textInput"
	case DataReceiverGeneratorType:
		return "dataReceiver"
	case ConstantSignalGeneratorType:
		return "constantSignal"
	default:
		return "wifi"
	}
}

type RunMode int

const (
	SenderMode = iota
	ReceiverMode
)

func (r RunMode) String() string {
	switch r {
	case SenderMode:
		return "senderMode"
	case ReceiverMode:
		return "receiverMode"
	default:
		return "unknown"
	}
}

type WebsocketConnectionInfo struct {
	Ip   string `json:"ip"`
	Port int    `json:"port"`
}

type ScreenExhibitionConfig struct {
	FilterUnderadable bool `json:"filterUnreadable"`
}
