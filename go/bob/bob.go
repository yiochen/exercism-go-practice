package bob // package name must match the package name in bob_test.go

import (
	"regexp"
	"unicode"
)

const testVersion = 2 // same as targetTestVersion

var isQuestion = regexp.MustCompile(`\?\s*$`)
var isEmpty = regexp.MustCompile(`^[\v\s\x{00a0}\x{2002}]*$`)

// Hey gives response as if bob was asked.
func Hey(input string) string {
	var containLowerCase = false
	var containUpperCase = false
	for _, rune := range input {
		if unicode.IsLetter(rune) {
			if unicode.ToLower(rune) == rune {
				containLowerCase = true
			}

			if unicode.ToUpper(rune) == rune {
				containUpperCase = true
			}
		}
	}

	if containUpperCase && !containLowerCase {
		return "Whoa, chill out!"
	}

	if isQuestion.MatchString(input) {
		return "Sure."
	}

	if isEmpty.MatchString(input) {
		return "Fine. Be that way!"
	}

	return "Whatever."
}
