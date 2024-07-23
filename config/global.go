package config

import "religare/models"

var AppLanguage models.Language

func init() {
	AppLanguage = models.English
}
