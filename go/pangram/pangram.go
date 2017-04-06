package pangram

import (
	"github.com/golang-collections/go-datastructures/set"
	"unicode"
)

const testVersion = 1

// IsPangram checks if a string is a pangram
func IsPangram(input string) bool {
	var hashSet = set.New()
	var isGerman = false
	for _, rune := range input {
		if len(string(rune)) > 2 {
			// not ascii
			return false
		}
		if unicode.IsLetter(rune) {
			hashSet.Add(unicode.ToUpper(rune), unicode.ToLower(rune))
		}
		if len(string(rune)) == 2 {
			isGerman = true
		}
	}
	if isGerman {
		return hashSet.Len() == 59 // Almighty Google told me this number
	}
	return hashSet.Len() == 52
}
