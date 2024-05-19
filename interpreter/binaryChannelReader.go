package interpreter

import (
	"fmt"
	"lazarus/models"
)

func ReadChannel(channel <-chan models.Binary) bool {
	for {
		if checkConsecutiveOnes(channel) {
			fmt.Println("Got a valid message")
		} else {
			fmt.Println("Invalid data")
		}
	}
}

func checkConsecutiveOnes(ch <-chan models.Binary) bool {
	consecutiveOnes := 0
	for binary := range ch {
		if binary == models.One {
			consecutiveOnes++
			if consecutiveOnes == 5 {
				return true
			}
		} else {
			return false
		}
	}
	return false
}
