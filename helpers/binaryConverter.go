package helpers

import (
	"fmt"
	"lazarus/models"
	"strconv"
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

func StringToBinaryString(value string, zfill int) string {
	var binaryString string

	for _, char := range value {
		binaryChar := strconv.FormatInt(int64(char), 2)
		binaryChar = fmt.Sprintf("%0*s", zfill, binaryChar)

		binaryString += binaryChar
	}

	return binaryString
}

func BinaryStringToBinaryData(value string) []models.Binary {
	binaryData := make([]models.Binary, len(value))

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
