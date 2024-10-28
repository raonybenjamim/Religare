/*
 * Religare - An Instrumental Trans Communication solution for communicating with paranormal entities
 *
 * Copyright (C) 2024 Raony Benjamim
 * Check the LICENSE.md file for more information regarding the code license
 */
package customTypes

type ReportLine struct {
	Character string
	IsValid   bool
	Attempts  int
}

type CalibrationConfig struct {
	CalibrationFileLocation string `json:"calibrationFileLocation"`
	EvaluationThreshold     int    `json:"evaluationThreshold"`
	ReportOutputPath        string `json:"reportOutputPath"`
}
