/*
 * Religare - An Instrumental Trans Communication solution for communicating with paranormal entities
 *
 * Copyright (C) 2024 Raony Benjamim
 * Check the LICENSE.md file for more information regarding the code license
 */

package converter

import (
	"log"
	"net"
	"religare/config"
	"religare/models"
	"time"
)

type DataSenderGenerator struct {
	channel           chan models.Binary
	internalGenerator SignalGenerator
	BufferSize        int
	connection        *net.UDPConn
}

func NewDataSenderGenerator(bufferSize int, generator SignalGenerator) *DataSenderGenerator {
	// evaluate global config for websocket communication
	if config.WebSocketConfig == nil {
		panic("data sender was created but no websocket configuration was provided. Please refer to the user manual")
	}

	writeAddr := &net.UDPAddr{IP: net.ParseIP("0.0.0.0"), Port: 2918} // we don't need to care about this address
	readAddr := &net.UDPAddr{IP: net.ParseIP(config.WebSocketConfig.Ip), Port: config.WebSocketConfig.Port}

	conn, err := net.DialUDP("udp", writeAddr, readAddr)

	if err != nil {
		log.Fatalf("Error dialing UDP: %v", err)
	}

	return &DataSenderGenerator{
		channel:           generator.GetChannel(),
		internalGenerator: generator,
		BufferSize:        bufferSize,
		connection:        conn,
	}
}

func (dsg *DataSenderGenerator) GenerateSignal() {
	// Generate signal to the channel (Start go routine for random generator)
	go dsg.internalGenerator.GenerateSignal()
	log.Printf("Starting data sender with the following config: %v \n", config.WebSocketConfig)

	// Send data every 100 ms
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	for range ticker.C {
		bytesQuantity := 4
		if _, ok := dsg.internalGenerator.(*ConstantSignalGenerator); ok {
			bytesQuantity = 1
		}
		// Read packages of 4 Bytes from the channel
		bytesToSend := readBytes(dsg.channel, bytesQuantity)
		// Send data to the websocket connection
		_, err := dsg.connection.Write(bytesToSend)

		if err != nil {
			log.Printf("Error while writing UDP data: %v\n", err)
		}
	}
}

func readBytes(channel chan models.Binary, quantity int) []byte {
	quantityInBits := quantity * models.ByteSize
	tempData := ""

	for i := 0; i < quantityInBits; i++ {
		bit := <-channel
		tempData += bit.String()
	}

	return []byte(tempData)
}

func (dsg *DataSenderGenerator) GetChannel() chan models.Binary {
	return dsg.channel
}

func (dsg *DataSenderGenerator) StopSignalGeneration() {
	close(dsg.channel)
}
