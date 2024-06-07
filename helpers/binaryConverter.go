package helpers

import (
	"fmt"
	"strconv"
)

func ConvertBinaryToInt(binaryString string) (int, error) {
	value, err := strconv.ParseInt(binaryString, 2, 64)

	if err != nil {
		return 0, fmt.Errorf("binary data is not convertable to int: %v", err)
	}

	return int(value), nil
}
