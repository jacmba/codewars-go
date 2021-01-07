package spinningwords

import "strings"

// SpinWords reverse words higher thatn 5 characteres in given string
func SpinWords(str string) string {
	words := strings.Split(str, " ")
	for i, w := range words {
		if len(w) >= 5 {
			chars := strings.Split(w, "")
			rev := []string{}
			for j := len(chars) - 1; j >= 0; j-- {
				rev = append(rev, chars[j])
			}
			words[i] = strings.Join(rev, "")
		}
	}
	return strings.Join(words, " ")
}
