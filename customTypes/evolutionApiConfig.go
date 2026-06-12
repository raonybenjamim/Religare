/*
 * Religare - An Instrumental Trans Communication solution for communicating with paranormal entities
 *
 * Copyright (C) 2024 Raony Benjamim
 * Check the LICENSE.md file for more information regarding the code license
 */
package customTypes

type EvolutionAPIConfig struct {
	Enabled      bool   `json:"enabled"`
	BaseURL      string `json:"baseUrl"`
	Instance     string `json:"instance"`
	APIKey       string `json:"apiKey"`
	TargetNumber string `json:"targetNumber"`
}

type SendTextPayload struct {
	Number           string   `json:"number"`
	Text             string   `json:"text"`
	Delay            int      `json:"delay"`
	LinkPreview      bool     `json:"linkPreview"`
	MentionsEveryOne bool     `json:"mentionsEveryOne"`
	Mentioned        []string `json:"mentioned,omitempty"`
}
