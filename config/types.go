/*
 * Religare - An Instrumental Trans Communication solution for communicating with paranormal entities
 *
 * Copyright (C) 2024 Raony Benjamim
 * Check the LICENSE.md file for more information regarding the code license
 */
package config

import (
	"encoding/json"
	"strings"
)

type CalibrationConfig struct {
	CalibrationFileLocation string `json:"calibrationFileLocation"`
	EvaluationThreshold     int    `json:"evaluationThreshold"`
	ReportOutputPath        string `json:"reportOutputPath"`
}

type ScreenExhibitionConfig struct {
	FilterUnderadable bool `json:"filterUnreadable"`
	ConstantReceiver  bool `json:"constantReceiver"` // only shows on screen when receives zeros
}

type WebsocketConnectionConfig struct {
	Ip          string `json:"ip"`
	Port        int    `json:"port"`
	MessageSize int    `json:"messageSize"`
}

type RunModeConfig int

const (
	SenderMode = iota
	ReceiverMode
	ExecuteLocalMode
)

func (r RunModeConfig) String() string {
	//TODO: Mudar essas strings e passar a usar o "RunMode" para melhor estruturar a execução na main
	switch r {
	case SenderMode:
		return "sender"
	case ReceiverMode:
		return "receiver"
	case ExecuteLocalMode:
		return "executeLocal"
	default:
		return "unknown"
	}
}

// Add the marshal and unmarshal functions for the RunMode type
func (r RunModeConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r *RunModeConfig) UnmarshalJSON(data []byte) error {
	// Here RunMode is unmarshalled as a string, so we need to convert it back to a RunMode
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	r.RunModeFromString(str)
	return nil
}

func (r *RunModeConfig) RunModeFromString(str string) {
	switch strings.ToLower(str) {
	case "sendermode":
		*r = SenderMode
	case "receivermode":
		*r = ReceiverMode
	default:
		*r = SenderMode
	}
}

type DataReadConfig struct {
	BypassValidation bool `json:"bypassValidation"`
}

type DataProduceConfig struct {
	ConstantSignal bool `json:"constantSignal"`
	MessageSize    int  `json:"messageSize"`
}
