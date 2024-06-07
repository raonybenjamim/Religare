package interpreter

import (
	"lazarus/models"
	"testing"
)

func TestCheckConsecutiveOnes(t *testing.T) {
	communicationChannel := make(chan models.Binary, 5)

	loadChannel(communicationChannel, models.One)

	// Check function should return true
	if !checkConsecutiveOnes(communicationChannel) {
		t.Error("Consecutive Ones should've returned true")
	}

	loadChannel(communicationChannel, models.Zero)
	if checkConsecutiveOnes(communicationChannel) {
		t.Error("Consecutive Zeros should've returned false")
	}
}

func loadChannel(ch chan models.Binary, value models.Binary) {
	for i := 0; i < 5; i++ {
		ch <- value
	}
}
