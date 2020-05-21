package function

func PlusFloat(params ...float64) float64 {
	var result float64 = 0
	for _, p := range params {
		result = result + p
	}

	return result
}
