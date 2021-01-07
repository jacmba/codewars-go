package duplicateencoder

import "strings"

func DuplicateEncode(word string) string {
	input := strings.ToUpper(word)
	occurrences := map[string]int{}

	var result string

	for i := 0; i < len(input); i++ {
		char := string(input[i])
		n, ok := occurrences[char]
		if ok {
			occurrences[char] = n + 1
		} else {
			occurrences[char] = 1
		}
	}

	for i := 0; i < len(input); i++ {
		char := string(input[i])
		n := occurrences[char]

		if n == 1 {
			result += "("
		} else {
			result += ")"
		}
	}

	return result
}
