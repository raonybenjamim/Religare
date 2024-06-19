package converter

import (
	"fmt"
	"lazarus/models"
	"os/exec"
	"regexp"
	"strconv"
)

type WifiSignalGenerator struct {
	channel    chan models.Binary
	BufferSize int
	threshold  float64
}

func NewWifiSignalGenerator(bufferSize int, threshhold float64) *WifiSignalGenerator {
	return &WifiSignalGenerator{
		channel:    make(chan models.Binary, bufferSize),
		BufferSize: bufferSize,
		threshold:  threshhold,
	}
}

func (wsg *WifiSignalGenerator) GetChannel() chan models.Binary {
	return wsg.channel
}

func (wsg *WifiSignalGenerator) GenerateSignal() {
	for {
		binary, err := convertToBinary(wsg.threshold)
		if err != nil {
			fmt.Println("Error:", err)
			close(wsg.channel)
			return
		}

		wsg.channel <- binary

	}
}

func getSignalStrength() (float64, error) {
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

func convertToBinary(threshold float64) (models.Binary, error) {
	strength, err := getSignalStrength()
	if err != nil {
		return 0, err
	}

	if strength > threshold {
		return 1, nil
	} else {
		return 0, nil
	}
}
