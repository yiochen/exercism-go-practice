// Package acronym provide utility to get acronym
package acronym

import (
	"fmt"
	"regexp"
	"strings"
)

const testVersion = 2

type StringSlice []string
type StringFunc func(s string) (result string)

func _map(s StringSlice, fn StringFunc) []string {
	var mapped []string
	for _, item := range s {
		fmt.Print(item)
		mapped = append(mapped, fn(item))

	}
	return mapped
}

func _getGroup(list [][]string) []string {
	var result []string
	for _, sublist := range list {
		if sublist[2] == "" {
			result = append(result, sublist[1])
		} else {
			result = append(result, sublist[2])
		}

	}
	return result
}

// Abbreviate get abbreviate of a string
func Abbreviate(input string) string {
	r := regexp.MustCompile(`(\b[a-zA-Z]|[a-z]([A-Z]))`)
	letters := r.FindAllStringSubmatch(input, -1)
	abbr := _map(_getGroup(letters), strings.ToUpper)
	return strings.Join(abbr, "")
}

//func main() {
//	Abbreviate("Portable Network Graphics")
//}
