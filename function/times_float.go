package function

import "math/big"

func TimesFloat(params ...float64) float64 {
	var result = big.NewFloat(params[0])

	rest := params[1:]

	for _, p := range rest {
		result = new(big.Float).Mul(result, big.NewFloat(p))
	}

	val, _ := result.Float64()

	return val
}
