package function

func DivInt(params ...int) int {
	result := 1
	for _, p := range params {
		result = result / p
	}

	return result
}
