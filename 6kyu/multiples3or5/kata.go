package multiples3or5

// Multiple3And5 return sum of multiples of 3 and 5 below given number
func Multiple3And5(n int) int {
	sum := 0
	for i := 1; i < n; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			sum += i
		}
	}
	return sum
}
