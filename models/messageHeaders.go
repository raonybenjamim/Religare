package models

const (
	ValidStart      string = "1111101010"
	ValidStartBits  int    = 10
	MessageTypeBits int    = 4
	MessageSizeBits int    = 10
	ChecksumBits    int    = 128
)

type messageType struct {
	Text      string
	TextFile  string
	Image     string
	ImageFile string
	Audio     string
	AudioFile string
}

var MessageType = messageType{
	Text:      "0001",
	TextFile:  "0010",
	Image:     "0011",
	ImageFile: "0100",
	Audio:     "0101",
	AudioFile: "0110",
}

type MessageHeaders struct {
	MessageType      string
	MessageSizeBytes int
	Checksum         string
}
