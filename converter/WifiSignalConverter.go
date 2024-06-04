package converter

import (
	"fmt"
	"lazarus/models"
	"os/exec"
	"regexp"
	"strconv"
)

func GetSignalStrength() (float64, error) {
	cmd := exec.Command("iwconfig")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return 0, err
	}

	re := regexp.MustCompile(`Link Quality=(\d+)/(\d+)`)
	matches := re.FindStringSubmatch(string(output))
	if matches == nil {
		return 0, fmt.Errorf("could not parse signal strength")
	}

	numerator, _ := strconv.Atoi(matches[1])
	denominator, _ := strconv.Atoi(matches[2])

	return float64(numerator) / float64(denominator), nil
}

func ConvertToBinary(threshold float64) (models.Binary, error) {
	strength, err := GetSignalStrength()
	if err != nil {
		return 0, err
	}

	if strength > threshold {
		return 1, nil
	} else {
		return 0, nil
	}
}

func AddBinaryData(threshold float64, ch chan<- models.Binary) {
	for {
		binary, err := ConvertToBinary(threshold)
		if err != nil {
			fmt.Println("Error:", err)
			close(ch)
			return
		}

		ch <- binary

	}
}
