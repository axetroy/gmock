package function

import "math/big"

func DivInt(params ...int64) int64 {
	var result = big.NewInt(params[0])

	rest := params[1:]

	for _, p := range rest {
		result = new(big.Int).Quo(result, big.NewInt(p))
	}

	return result.Int64()
}