package util

import (
	"math/rand"
	"time"
)

func RandomString(length int) (result string) {
	rand.Seed(time.Now().Unix())

	ranStr := make([]byte, length)

	// Generating Random string
	for i := 0; i < length; i++ {
		ranStr[i] = byte(65 + rand.Intn(25))
	}

	result = string(ranStr)
	return
}
