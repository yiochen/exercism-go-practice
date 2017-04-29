package triangle

import "math"

const testVersion = 3

// Code this function.
func KindFromSides(a, b, c float64) Kind {
	if math.IsNaN(a)  || math.IsNaN(b) || math.IsNaN(c) {
		return NaT
	}
	if math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) {
		return NaT
	}
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}
	halfOfSum := (a + b + c) / 2
	if a > halfOfSum || b > halfOfSum || c > halfOfSum {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || b == c || a == c {
		return Iso
	}

	return Sca
}

// Notice it returns this type.  Pick something suitable.
type Kind string

// Pick values for the following identifiers used by the test program.
var NaT Kind = "nat" // not a triangle
var Equ Kind = "equ" // equilateral
var Iso Kind = "iso" // isosceles
var Sca Kind = "sca" // scalene

// Organize your code for readability.
