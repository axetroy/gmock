package function

import (
	"math/rand"
	"time"
)

func RandomStr(length int, base ...string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	if len(base) > 0 {
		letterRunes = []rune(base[0])
	}

	b := make([]rune, length)

	for i := range b {
		b[i] = letterRunes[r.Intn(len(letterRunes))]
	}
	return string(b)
}
