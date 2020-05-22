package function

import (
	"fmt"
	"math/rand"
	"time"
)

func RangeFloat(min float64, max float64) float64 {
	if min > max {
		panic(fmt.Errorf("rangeFloat: the min `%f` is larger then max `%f`", min, max))
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	n := r.Float64() * max

	if n < min {
		n = n + min
	}

	return min
}
