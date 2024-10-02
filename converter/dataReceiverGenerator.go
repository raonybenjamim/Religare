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
	"religare/helpers"
	"religare/models"
)

type DataReceiverGenerator struct {
	channel    chan models.Binary
	BufferSize int
	connection *net.UDPConn
}

func NewDataReceiverGenerator(bufferSize int) *DataReceiverGenerator {
	// evaluate global config for websocket communication
	if config.WebSocketConfig == nil {
		panic("data sender was created but no websocket configuration was provided. Please refer to the user manual")
	}

	address := &net.UDPAddr{IP: net.ParseIP(config.WebSocketConfig.Ip), Port: config.WebSocketConfig.Port}

	udpConnection, err := net.ListenUDP("udp", address)

	if err != nil {
		panic("Not able to start the server")
	}

	return &DataReceiverGenerator{
		channel:    make(chan models.Binary, bufferSize),
		BufferSize: bufferSize,
		connection: udpConnection,
	}
}

func (drg *DataReceiverGenerator) GenerateSignal() {
	// Read the first 32 bits of data from the connection
	readBuffer := make([]byte, 4*models.ByteSize)

	log.Printf("Starting data sender with the following config: %v \n", config.WebSocketConfig)
	for {
		n, _, err := drg.connection.ReadFromUDP(readBuffer)

		if err != nil {
			log.Printf("Got Error while reading connection: %v\n", err)
		}

		//TODO: may be interesting watching what may be the result of using stringToBinaryString first here
		readData := string(readBuffer[:n])
		// convert it to models.Binary
		binaryData := helpers.BinaryStringToBinaryData(readData)
		// send the data to the data channel
		for _, bit := range binaryData {
			drg.channel <- bit
		}
	}

}

func (drg *DataReceiverGenerator) GetChannel() chan models.Binary {
	return drg.channel
}

func (drg *DataReceiverGenerator) StopSignalGeneration() {
	close(drg.channel)
}
