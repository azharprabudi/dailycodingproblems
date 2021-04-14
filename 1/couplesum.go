package couplesum

func CoupleSum(arr []int, k int) [2]int {
	dicts := map[int]int{}
	for _, numb := range arr {
		_, exists := dicts[numb]
		if exists {
			return [2]int{dicts[numb], numb}
		}

		dicts[k-numb] = numb
	}

	return [2]int{}
}
