package models

const (
	ConverterChannelSize = 100
	ByteSize             = 8
)

type generatorType struct {
	Random    string
	Wifi      string
	TextInput string
}

var GeneratorType = generatorType{
	Random:    "random",
	Wifi:      "wifi",
	TextInput: "textInput",
}
