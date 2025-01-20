package app

import (
	"fmt"
	"regexp"
	"unicode"
)

func countAlphanumeric(input string) int {
	count := 0

	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			count++
		}
	}

	return count
}

func validateRegex(input string, pattern string) bool {
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Printf("Invalid regex pattern: %v\n", err)
		return false
	}

	return re.MatchString(input)
}
