package function

import (
	"fmt"
	"math/rand"
	"time"
)

func RangeInt(min int, max int) int {
	if min > max {
		panic(fmt.Errorf("rangeFloat: the min `%d` is larger then max `%d`", min, max))
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	n := r.Intn(max)

	if n < min {
		n = n + min
	}

	return min
}
