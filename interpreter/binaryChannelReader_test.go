package interpreter

import (
	"lazarus/models"
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

var mockChecksum = createMockChecksum(models.ChecksumBits)
var mockMessageSize = zfill(strconv.FormatInt(20, 2), models.MessageSizeBits)
var validHeadersForTextMessage = models.MessageType.Text + mockChecksum + mockMessageSize

func TestGetHeaders(t *testing.T) {
	communicationChannel := make(chan models.Binary, len(validHeadersForTextMessage))

	channelReader := ChannelReader{
		Channel: communicationChannel,
	}

	loadChannel(communicationChannel, validHeadersForTextMessage)

	headers, err := channelReader.readHeadersFromChannel()

	// Check function should return true
	if err != nil {
		t.Fatalf("Error while reading headers %v", err)
	}

	if headers.Checksum != mockChecksum {
		t.Errorf("Header Checksum was not equal. Expected: %v, Got: %v", mockChecksum, headers.Checksum)
	}

	if headers.MessageSizeBytes != 20 {
		t.Errorf("Header Message Size was not equal. Expected: %v, Got: %v", mockMessageSize, headers.MessageSizeBytes)
	}

	if headers.MessageType != models.MessageType.Text {
		t.Errorf("Header Message Type was not equal. Expected %v, Got: %v", models.MessageType.Text, headers.MessageType)
	}
}

func TestBinaryToText(t *testing.T) {
	tests := []struct {
		binaryString string
		expectedText string
		expectError  bool
	}{
		{"0100100001100101011011000110110001101111", "Hello", false},
		{"0110100001100101011011000110110001101111", "hello", false},
		{"010000010100001001000011", "ABC", false},
		{"001100010011001000110011", "123", false},
		{"", "", false},             // empty string should return an empty string
		{"00000000", "\x00", false}, // single null byte
		{"0100100", "", true},       // invalid length (not a multiple of 8)
		{"0100100G", "", true},      // invalid character in binary string
	}

	communicationChannel := make(chan models.Binary, len(validHeadersForTextMessage))

	channelReader := ChannelReader{
		Channel: communicationChannel,
	}

	for _, test := range tests {
		result, err := channelReader.binaryToText(test.binaryString)
		if (err != nil) != test.expectError {
			t.Errorf("binaryToText(%q) error = %v, expected error = %v", test.binaryString, err, test.expectError)
			continue
		}
		if result != test.expectedText {
			t.Errorf("binaryToText(%q) = %q, want %q", test.binaryString, result, test.expectedText)
		}
	}
}

func TestValidMessageIsCorrectlyChecked(t *testing.T) {
	tests := []struct {
		expectedBinaryChecksum string
		textMessage            string
		expectedError          bool
	}{
		{
			"0011011101100010001100100110000100110011001101000110001101100001011000100011001101100100011000110011001000110011011000110011100000110000011000100011100101100110001101000110010001100010011001100011100100110011011000110110001100110000011000100110010001100101",
			"Correct Testing message",
			true},
		{
			"0011011101100010001100100110000100110011001101100110001101100001011000100011001101100100011000110011001000110011011000110011100000110000011000100011100101100110001101000110010001100010011001100011100100110011011000110110001100110000011000100110010001100101",
			"This message will fail",
			false},
	}

	channelReader := ChannelReader{
		Channel: make(chan models.Binary),
	}

	for _, test := range tests {
		if channelReader.isValidStringMessage(test.expectedBinaryChecksum, test.textMessage) != test.expectedError {
			t.Errorf("Error while checking vality for message: %v", test.textMessage)
		}
	}
}

func loadChannel(ch chan models.Binary, value string) {
	for _, bit := range value {
		switch bit {
		case '0':
			ch <- models.Zero
		case '1':
			ch <- models.One
		}
	}
}

func createMockChecksum(size int) string {
	checksum := make([]byte, size) // Create a byte slice of the given size

	// Generate random 0s and 1s
	for i := 0; i < size; i++ {
		// Generate a random number between 0 and 1
		bit := rand.Intn(2)

		// Convert the random number to a byte (0 or 1) and store it in the checksum
		checksum[i] = byte('0' + bit)
	}

	return string(checksum) // Convert byte slice to string and return
}

func zfill(str string, length int) string {
	if len(str) >= length {
		return str // If the string is already equal or longer than the desired length, return it as is
	}
	zerosNeeded := length - len(str)          // Calculate the number of zeros needed
	zeros := strings.Repeat("0", zerosNeeded) // Create a string of zeros
	return zeros + str                        // Concatenate the zeros with the original string
}
