package function

import (
	"math/rand"
	"time"
)

func RangeInt(min int, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	n := r.Intn(max)

	if n < min {
		n = n + min
	}

	return min
}
