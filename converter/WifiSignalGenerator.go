/*
 * Religare - An Instrumental Trans Communication solution for communicating with paranormal entities
 *
 * Copyright (C) 2024 Raony Benjamim
 * Check the LICENSE.md file for more information regarding the code license
 */
package converter

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"religare/models"
	"runtime"
	"strconv"
	"strings"
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

func (wsg *WifiSignalGenerator) StopSignalGeneration() {
	close(wsg.channel)
}

func getSignalStrength() (float64, error) {
	switch runtime.GOOS {
	case "windows":
		return getWifiSignalStrengthWindows()
	case "linux":
		return getWifiSignalStrengthLinux()
	default:
		return 0, fmt.Errorf("unsupported platform")
	}
}

func getWifiSignalStrengthWindows() (float64, error) {
	cmd := exec.Command("netsh", "wlan", "show", "interfaces")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return 0, err
	}
	output := out.String()
	signalStrength, err := parseSignalStrengthWindows(output)
	if err != nil {
		return 0, err
	}
	return float64(signalStrength) / 100, nil
}

func parseSignalStrengthWindows(output string) (int, error) {
	re := regexp.MustCompile(`:\s*(\d+)\s*%`)
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if match != nil && len(match) == 2 {
			return strconv.Atoi(match[1])
		}
	}
	return 0, fmt.Errorf("signal strength not found")
}

func getWifiSignalStrengthLinux() (float64, error) {
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
		return models.One, nil
	} else {
		return models.Zero, nil
	}
}
