package function

func MinusInt(params ...int) int {
	result := 0
	for _, p := range params {
		result = result - p
	}

	return result
}
