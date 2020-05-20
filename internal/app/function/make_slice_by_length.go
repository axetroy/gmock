package function

func MakeSliceByLength(length int) []int {
	arr := make([]int, 0)
	for i := 0; i < length; i++ {
		arr = append(arr, i)
	}

	return arr
}
