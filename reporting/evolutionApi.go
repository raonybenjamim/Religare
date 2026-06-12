/*
 * Religare - An Instrumental Trans Communication solution for communicating with paranormal entities
 *
 * Copyright (C) 2024 Raony Benjamim
 * Check the LICENSE.md file for more information regarding the code license
 */
package reporting

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"religare/config"
	"religare/customTypes"
	"time"
)

var activityChannel = make(chan bool, 10)

// RegisterActivity signals that a positive result was obtained, resetting the heartbeat timer.
func RegisterActivity() {
	select {
	case activityChannel <- true:
	default:
	}
}

// SendTextMessage sends a message through the Evolution API.
func SendTextMessage(text string) error {
	if config.EvolutionAPIConfig == nil || !config.EvolutionAPIConfig.Enabled {
		return nil
	}

	cfg := config.EvolutionAPIConfig

	payload := customTypes.SendTextPayload{
		Number:           cfg.TargetNumber,
		Text:             text,
		Delay:            1000,
		LinkPreview:      true,
		MentionsEveryOne: false,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal evolution API payload: %v", err)
	}

	url := fmt.Sprintf("%s/message/sendText/%s", cfg.BaseURL, cfg.Instance)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apikey", cfg.APIKey)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send message via evolution API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("evolution API returned status code: %d", resp.StatusCode)
	}

	return nil
}

// StartDailyHeartbeat checks for inactivity and sends a heartbeat message if no results were obtained.
func StartDailyHeartbeat(runMode string) {
	if config.EvolutionAPIConfig == nil || !config.EvolutionAPIConfig.Enabled {
		return
	}

	go func() {
		for {
			// Wait for 24 hours or until an activity is registered
			select {
			case <-time.After(24 * time.Hour):
				message := fmt.Sprintf("Estou bem, mas não tive resultados ainda. Modo de execução: %s", runMode)
				err := SendTextMessage(message)
				if err != nil {
					fmt.Printf("Error sending daily heartbeat: %v\n", err)
				}
			case <-activityChannel:
				// Activity registered, loop restarts and we wait another 24h
			}
		}
	}()
}
