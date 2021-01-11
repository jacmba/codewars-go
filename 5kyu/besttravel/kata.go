package besttravel

func getCombinations(k int, ls []int) []int {
	if k == 1 {
		return ls
	}

	result := []int{}

	for i := 0; i < len(ls)-k+1; i++ {
		arr := getCombinations(k-1, ls[i+1:])
		for _, v := range arr {
			result = append(result, ls[i]+v)
		}
	}

	return result
}

// ChooseBestSum return best distance
func ChooseBestSum(t, k int, ls []int) int {
	best := -1

	combs := getCombinations(k, ls)
	for _, v := range combs {
		if v > best && v <= t {
			best = v
		}
	}

	return best
}
