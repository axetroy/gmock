package function

import (
	"errors"
	"math/big"
	"strconv"
)

func Div(params ...interface{}) float64 {
	result := big.NewFloat(0)

	for index, p := range params {
		switch v := p.(type) {
		case int:
			if index == 0 {
				result = result.Add(result, big.NewFloat(float64(v)))
			} else {
				result = new(big.Float).Quo(result, big.NewFloat(float64(v)))
			}
		case uint:
			if index == 0 {
				result = result.Add(result, big.NewFloat(float64(v)))
			} else {
				result = new(big.Float).Quo(result, big.NewFloat(float64(v)))
			}
		case int8:
			if index == 0 {
				result = result.Add(result, big.NewFloat(float64(v)))
			} else {
				result = new(big.Float).Quo(result, big.NewFloat(float64(v)))
			}
		case uint8:
			if index == 0 {
				result = result.Add(result, big.NewFloat(float64(v)))
			} else {
				result = new(big.Float).Quo(result, big.NewFloat(float64(v)))
			}
		case int16:
			if index == 0 {
				result = result.Add(result, big.NewFloat(float64(v)))
			} else {
				result = new(big.Float).Quo(result, big.NewFloat(float64(v)))
			}
		case uint16:
			if index == 0 {
				result = result.Add(result, big.NewFloat(float64(v)))
			} else {
				result = new(big.Float).Quo(result, big.NewFloat(float64(v)))
			}
		case int32:
			if index == 0 {
				result = result.Add(result, big.NewFloat(float64(v)))
			} else {
				result = new(big.Float).Quo(result, big.NewFloat(float64(v)))
			}
		case uint32:
			if index == 0 {
				result = result.Add(result, big.NewFloat(float64(v)))
			} else {
				result = new(big.Float).Quo(result, big.NewFloat(float64(v)))
			}
		case int64:
			if index == 0 {
				result = result.Add(result, big.NewFloat(float64(v)))
			} else {
				result = new(big.Float).Quo(result, big.NewFloat(float64(v)))
			}
		case uint64:
			if index == 0 {
				result = result.Add(result, big.NewFloat(float64(v)))
			} else {
				result = new(big.Float).Quo(result, big.NewFloat(float64(v)))
			}
		case float32:
			if index == 0 {
				result = result.Add(result, big.NewFloat(float64(v)))
			} else {
				result = new(big.Float).Quo(result, big.NewFloat(float64(v)))
			}
		case float64:
			if index == 0 {
				result = result.Add(result, big.NewFloat(float64(v)))
			} else {
				result = new(big.Float).Quo(result, big.NewFloat(float64(v)))
			}
		case string:
			if f, err := strconv.ParseFloat(v, 10); err != nil {
				panic(err)
			} else {
				if index == 0 {
					result = result.Add(result, big.NewFloat(float64(f)))
				} else {
					result = new(big.Float).Quo(result, big.NewFloat(f))
				}
			}
		default:
			panic(errors.New("Invalid type for 'Div' function"))
		}
	}

	val, _ := result.Float64()

	return val
}
