// Package raindrops
package raindrops

import (
	"bytes"
	"fmt"
)

const testVersion = 2

// Convert convert the number to a string
func Convert(num int) string {
	var buffer bytes.Buffer

	if num%3 == 0 {
		buffer.WriteString("Pling")
	}

	if num%5 == 0 {
		buffer.WriteString("Plang")
	}

	if num%7 == 0 {
		buffer.WriteString("Plong")
	}

	if buffer.Len() == 0 {
		return fmt.Sprint(num)
	}

	return buffer.String()

}
