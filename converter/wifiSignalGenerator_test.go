/*
 * Religare - An Instrumental Trans Communication solution for communicating with paranormal entities
 *
 * Copyright (C) 2024 Raony Benjamim
 * Check the LICENSE.md file for more information regarding the code license
 */
package converter

import (
	"lazarus/models"
	"testing"
)

func TestValueFromWifiSignal(t *testing.T) {
	signalStrength, err := getSignalStrength()

	// Check if there was an err
	if err != nil {
		t.Fatalf("GetSignalStrength() returned an error: %v", err)
	}

	// Check if the signalStrength is within the expected range
	if signalStrength < 0.0 || signalStrength > 1.0 {
		t.Errorf("GetSignalStrength() returned an out of range value: %f", signalStrength)
	}
}

func TestBinaryConvertFunction(t *testing.T) {
	threshold := 0.5

	binaryData, err := convertToBinary(threshold)

	if err != nil {
		t.Fatalf("ConvertToBinary returned an error: %v", err)
	}

	if binaryData != models.Zero && binaryData != models.One {
		t.Errorf("binaryData is neither Zero nor One: %v", binaryData)
	}
}
