/*
 * Religare - An Instrumental Trans Communication solution for communicating with paranormal entities
 *
 * Copyright (C) 2024 Raony Benjamim
 * Check the LICENSE.md file for more information regarding the code license
 */
package interpreter

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"religare/config"
	"religare/customTypes"
	"religare/helpers"
	"religare/models"
	"strings"
	"syscall"
	"time"
)

type BinaryDataBypassReader struct {
	Channel <-chan models.Binary
}

type ConstantReceiverStatistics struct {
	ZeroCount int
	OneCount  int
	Total     int
}

func (reader *BinaryDataBypassReader) GetChannel() <-chan models.Binary {
	return reader.Channel
}

func (reader *BinaryDataBypassReader) ReadChannel() {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	statistics := ConstantReceiverStatistics{}
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func(sig chan os.Signal, screenExhibitionConfig *customTypes.ScreenExhibitionConfig) {
		<-sig
		if screenExhibitionConfig.ConstantReceiver {
			fmt.Println("received signal, writing statistics to file")
			writeStatisticsToFile(&statistics)
		}
		os.Exit(0)
	}(sigs, config.ScreenExhibitionConfig)

	for range ticker.C {
		binaryCharacter := helpers.GetDataFromBinaryChannel(reader.Channel, 4*models.ByteSize)

		if config.ScreenExhibitionConfig.ConstantReceiver {
			checkForZeros(binaryCharacter, &statistics)
		} else {
			textCharacter, err := helpers.BinaryStringToString(binaryCharacter)

			if err != nil {
				fmt.Println("error while reading binary data form channel: " + err.Error())
			}
			showCharacters(textCharacter)
		}
	}
}

func writeStatisticsToFile(statistics *ConstantReceiverStatistics) {
	payload, err := json.Marshal(statistics)
	if err != nil {
		fmt.Println("error while marshalling statistics: " + err.Error())
	}
	helpers.WriteStringToFile(string(payload))
}

func checkForZeros(binaryCharacter string, statistics *ConstantReceiverStatistics) {
	zeroes := strings.Count(binaryCharacter, "0")
	if zeroes > 0 {
		statistics.ZeroCount += zeroes
		fmt.Println("===0===")
	} else {
		statistics.OneCount += strings.Count(binaryCharacter, "1")
	}
	statistics.Total += len(binaryCharacter)
}

func showCharacters(textCharacters string) {
	if !config.ScreenExhibitionConfig.FilterUnderadable {
		fmt.Print(textCharacters)
		return
	}

	fmt.Print(helpers.FilterUnderadableCharacters(textCharacters))
}
