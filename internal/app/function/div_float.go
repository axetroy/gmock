package function

func DivFloat(params ...float64) float64 {
	var result float64 = 1
	for _, p := range params {
		result = result / p
	}

	return result
}
