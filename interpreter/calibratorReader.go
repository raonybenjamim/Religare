/*
 * Religare - An Instrumental Trans Communication solution for communicating with paranormal entities
 *
 * Copyright (C) 2024 Raony Benjamim
 * Check the LICENSE.md file for more information regarding the code license
 */
package interpreter

import (
	"fmt"
	"religare/config"
	"religare/customTypes"
	"religare/helpers"
	"religare/models"
	"strings"
)

type Calibrator struct {
	Channel            <-chan models.Binary
	CalibrationChannel chan string
}

func (calibrator *Calibrator) GetChannel() <-chan models.Binary {
	return calibrator.Channel
}

func (calibrator *Calibrator) ReadChannel() {

	// maybe add some ticker
	for {
		binaryCharacter := helpers.GetDataFromBinaryChannel(calibrator.Channel, 4*models.ByteSize)

		textCharacter, err := helpers.BinaryStringToString(binaryCharacter)

		if err != nil {
			fmt.Println("error while reading binary data form channel: " + err.Error())
		}

		filteredData := helpers.FilterUnderadableCharacters(textCharacter)

		if filteredData != "" {
			calibrator.CalibrationChannel <- filteredData
		}
	}
}

func (calibrator *Calibrator) Calibrate() {
	dataToCalibrate, err := helpers.ReadCalibrationFile()
	if err != nil {
		panic(err)
	}

	evaluationThreshold := config.CalibrationConfig.EvaluationThreshold
	if evaluationThreshold == 0 {
		panic("cannot work with Evaluation Threshold zero. Please configure it in the config file")
	}

	reportData := []customTypes.ReportLine{}
	// for each character, read characters until threshold
	for _, char := range dataToCalibrate {
		for i := 0; i < evaluationThreshold; i++ {
			receivedData, ok := <-calibrator.CalibrationChannel
			if !ok {
				panic("error reading calibration channel")
			}
			if strings.EqualFold(string(char), receivedData) {
				reportData = append(reportData, customTypes.ReportLine{
					Character: receivedData,
					IsValid:   true,
					Attempts:  i,
				})
				i = evaluationThreshold
			}

			if i == evaluationThreshold {
				reportData = append(reportData, customTypes.ReportLine{
					Character: string(char),
					IsValid:   false,
					Attempts:  evaluationThreshold,
				})
			}
		}
	}
	err = helpers.WriteReportFile(reportData)
	if err != nil {
		panic(err)
	}
}
