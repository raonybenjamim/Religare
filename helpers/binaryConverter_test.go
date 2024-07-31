/*
 * Religare - An Instrumental Trans Communication solution for communicating with paranormal entities
 *
 * Copyright (C) 2024 Raony Benjamim
 * Check the LICENSE.md file for more information regarding the code license
 */
package helpers

import (
	"testing"
)

func TestBinaryToText(t *testing.T) {
	tests := []struct {
		binaryString string
		expectedText string
		expectError  bool
	}{
		{"0100100001100101011011000110110001101111", "Hello", false},
		{"0110100001100101011011000110110001101111", "hello", false},
		{"010000010100001001000011", "ABC", false},
		{"001100010011001000110011", "123", false},
		{"", "", false},             // empty string should return an empty string
		{"00000000", "\x00", false}, // single null byte
		{"0100100", "", true},       // invalid length (not a multiple of 8)
		{"0100100G", "", true},      // invalid character in binary string
	}

	for _, test := range tests {
		result, err := BinaryStringToString(test.binaryString)
		if (err != nil) != test.expectError {
			t.Errorf("binaryToText(%q) error = %v, expected error = %v", test.binaryString, err, test.expectError)
			continue
		}
		if result != test.expectedText {
			t.Errorf("binaryToText(%q) = %q, want %q", test.binaryString, result, test.expectedText)
		}
	}
}
