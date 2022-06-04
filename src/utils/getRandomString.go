package utils

import (
	"math/rand"
	"os"
	"strings"
	"time"
)

func GetRandomString() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	var b strings.Builder
	for i := 0; i < 5; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return os.Getenv("domen") + b.String()

}
