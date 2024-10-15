package utils

import (
	"math/rand"
	"time"
)

func RandString(charset string, strlen int) string {
	rand.New(rand.NewSource(time.Now().Unix()))
	result := make([]byte, strlen)

	for i := 0; i < strlen; i++ {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
