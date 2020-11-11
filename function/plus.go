package function

import (
	"errors"
	"math/big"
	"strconv"
)

func Plus(params ...interface{}) float64 {
	result := big.NewFloat(0)

	for _, p := range params {
		switch v := p.(type) {
		case int:
			result = new(big.Float).Add(result, big.NewFloat(float64(v)))
		case uint:
			result = new(big.Float).Add(result, big.NewFloat(float64(v)))
		case int8:
			result = new(big.Float).Add(result, big.NewFloat(float64(v)))
		case uint8:
			result = new(big.Float).Add(result, big.NewFloat(float64(v)))
		case int16:
			result = new(big.Float).Add(result, big.NewFloat(float64(v)))
		case uint16:
			result = new(big.Float).Add(result, big.NewFloat(float64(v)))
		case int32:
			result = new(big.Float).Add(result, big.NewFloat(float64(v)))
		case uint32:
			result = new(big.Float).Add(result, big.NewFloat(float64(v)))
		case int64:
			result = new(big.Float).Add(result, big.NewFloat(float64(v)))
		case uint64:
			result = new(big.Float).Add(result, big.NewFloat(float64(v)))
		case float32:
			result = new(big.Float).Add(result, big.NewFloat(float64(v)))
		case float64:
			result = new(big.Float).Add(result, big.NewFloat(v))
		case string:
			if f, err := strconv.ParseFloat(v, 10); err != nil {
				panic(err)
			} else {
				result = new(big.Float).Add(result, big.NewFloat(f))
			}
		default:
			panic(errors.New("Invalid type for 'Plus' function"))
		}
	}

	val, _ := result.Float64()

	return val
}
