package commondenominators

import (
	"fmt"
)

// ConvertFracts returns common factor fractions
func ConvertFracts(a [][]int) string {
	var result string
	var b [][]int

	for _, n := range a {
		min := n[0]
		if n[1] < min {
			min = n[1]
		}

		minCommon := 1

		for i := min; i > 1 && minCommon == 1; i-- {
			if n[0]%i == 0 && n[1]%i == 0 {
				minCommon = i
			}
		}

		b = append(b, []int{n[0] / minCommon, n[1] / minCommon})
	}

	max := b[0][1]

	for i := 1; i < len(b); i++ {
		if b[i][1] > max {
			max = b[i][1]
		}
	}

	maxCommon := true
	common := max

	for _, n := range b {
		if max%n[1] != 0 {
			maxCommon = false
			break
		}
	}

	if !maxCommon {
		for i := max * 2; common == max; i += max {
			isCommon := true
			for _, n := range b {
				if i%n[1] != 0 {
					isCommon = false
					break
				}
			}
			if isCommon {
				common = i
			}
		}
	}

	for _, n := range b {
		d := common / n[1]
		c := []int{n[0] * d, n[1] * d}
		result += fmt.Sprintf("(%d,%d)", c[0], c[1])
	}

	return result
}
