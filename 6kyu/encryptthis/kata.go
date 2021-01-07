package encryptthis

import (
	"fmt"
	"strings"
)

// EncryptThis encrypt given text string
func EncryptThis(text string) string {
	result := []string{}

	words := strings.Split(text, " ")
	for _, w := range words {
		runes := []rune(w)
		output := ""

		if len(runes) > 0 {
			output += fmt.Sprint(int(runes[0]))

			if len(runes) > 1 {
				if len(runes) > 2 {
					aux := runes[1]
					runes[1] = runes[len(runes)-1]
					runes[len(runes)-1] = aux
				}

				output += string(runes[1:])
			}
		}

		result = append(result, output)
	}

	return strings.Join(result, " ")
}
