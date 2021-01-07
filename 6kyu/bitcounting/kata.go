package bitcounting

import (
	"fmt"
)

// CountBits return number of 1s in binary representation of given number
func CountBits(n uint) int {
	bin := fmt.Sprintf("%b", n)
	count := 0

	for _, i := range []rune(bin) {
		if i == '1' {
			count++
		}
	}

	return count
}
