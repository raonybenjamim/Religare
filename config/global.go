/*
 * Religare - An Instrumental Trans Communication solution for communicating with paranormal entities
 *
 * Copyright (C) 2024 Raony Benjamim
 * Check the LICENSE.md file for more information regarding the code license
 */
package config

import (
	"religare/customTypes"
)

var AppLanguage customTypes.Language
var RunMode RunModeConfig
var GeneratorType customTypes.GeneratorType
var WebSocketConnection *WebsocketConnectionConfig
var ScreenExhibition *ScreenExhibitionConfig
var Calibration *CalibrationConfig
var DataRead *DataReadConfig
var DataProduce *DataProduceConfig
