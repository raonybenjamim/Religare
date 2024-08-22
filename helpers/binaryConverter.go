/*
 * Religare - An Instrumental Trans Communication solution for communicating with paranormal entities
 *
 * Copyright (C) 2024 Raony Benjamim
 * Check the LICENSE.md file for more information regarding the code license
 */
package helpers

import (
	"errors"
	"fmt"
	"religare/models"
	"strconv"
	"strings"
)

func ConvertBinaryToInt(binaryString string) (int, error) {
	value, err := strconv.ParseInt(binaryString, 2, 64)

	if err != nil {
		return 0, fmt.Errorf("binary data is not convertable to int: %v", err)
	}

	return int(value), nil
}

func IntToBinaryString(value int, zfill int) string {
	return fmt.Sprintf("%0*s", zfill, strconv.FormatInt(int64(value), 2))
}

/*
*
In order to allow utf-8 support we need to work with 32 bits runes
This function will encode the runes into bytes and then to binary string
*
*/
func StringToBinaryString(value string) string {
	var binaryString string

	for _, char := range value {
		//convert the rune to utf-8 bytes
		bytes := []byte(string(char))
		for _, utf8_byte := range bytes {
			binaryString += fmt.Sprintf("%08b", utf8_byte)
		}
	}

	return binaryString
}

func HexTo4BitBinary(hexString string) (string, error) {
	var binaryString strings.Builder

	for _, char := range hexString {
		// Convert the hex character to an integer
		value, err := strconv.ParseInt(string(char), 16, 64)
		if err != nil {
			return "", err
		}
		// Convert the integer to a 4-bit binary string
		binaryChar := fmt.Sprintf("%04b", value)
		binaryString.WriteString(binaryChar)
	}

	return binaryString.String(), nil
}

func BinaryStringToBinaryData(value string) []models.Binary {
	var binaryData []models.Binary

	for _, bitChar := range value {

		switch bitChar {
		case '0':
			binaryData = append(binaryData, models.Zero)
		case '1':
			binaryData = append(binaryData, models.One)
		}
	}

	return binaryData
}

func BinaryStringToString(binaryString string) (string, error) {

	if len(binaryString)%models.ByteSize != 0 {
		return "", errors.New("binary string length is not a multiple of 8")
	}

	var bytes []byte

	for i := 0; i < len(binaryString); i += models.ByteSize {
		// Get the next 8 bits
		byteString := binaryString[i : i+models.ByteSize]

		// Convert the 8-bit binary string to a decimal (base 10) integer
		byteValue, err := strconv.ParseUint(byteString, 2, models.ByteSize)
		if err != nil {
			return "", err
		}

		bytes = append(bytes, byte(byteValue))
	}

	return string(bytes), nil
}

func BinaryStringToHexString(binaryString string) (string, error) {
	if len(binaryString)%4 != 0 {
		return "", errors.New("binary string length is not a multiple of 4")
	}

	var hexString string

	for i := 0; i < len(binaryString); i += 4 {
		// Get the next 4 bits
		bitChunk := binaryString[i : i+4]

		// Convert the 4-bit binary string to a decimal (base 10) integer
		decimalValue, err := strconv.ParseInt(bitChunk, 2, 64)
		if err != nil {
			return "", err
		}

		// Convert the integer to a corresponding hexadecimal character
		hexString += strconv.FormatInt(decimalValue, 16)
	}

	return hexString, nil
}
