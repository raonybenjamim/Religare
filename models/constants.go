package models

const (
	ConverterChannelSize = 100
	ByteSize             = 8
	WifiThreshold        = 0.75
)

type Language int

const (
	Portuguese Language = iota
	English
)

func (l Language) String() string {
	switch l {
	case Portuguese:
		return "Portuguese"

	case English:
		return "English"

	default:
		return "unknown"
	}
}
