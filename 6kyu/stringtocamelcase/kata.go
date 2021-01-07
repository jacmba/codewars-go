package stringtocamelcase

import "strings"

// ToCamelCase return given string converted to camel case
func ToCamelCase(s string) string {
	var result string
	var input = strings.Split(s, "")
	var last string

	for _, c := range input {
		if c != "-" && c != "_" {
			if last == "-" || last == "_" {
				result += strings.ToUpper(c)
			} else {
				result += c
			}
		}
		last = c
	}
	return result
}
