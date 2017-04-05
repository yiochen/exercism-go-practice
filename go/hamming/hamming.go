// Package hamming include utiltiies for calculating the hamming distance.
package hamming

import "errors"

const testVersion = 5

// Distance calculate the hamming distance of two strings. return -1 and throw error if strings are of different lengths
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("strings should be the same length")
	}
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}

	return count, nil
}
