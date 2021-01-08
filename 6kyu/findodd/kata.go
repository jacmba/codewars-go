package findodd

// FindOdd return the number that appears an odd number of times
func FindOdd(seq []int) int {
	ocurrences := map[int]int{}

	for _, n := range seq {
		ocurrences[n]++
	}

	for k, v := range ocurrences {
		if v%2 != 0 {
			return k
		}
	}

	return 0
}
