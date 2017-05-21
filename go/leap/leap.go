//Package leap provide functionality for checking leap year
package leap

const testVersion = 3

// IsLeapYear check if a year is a leap year
func IsLeapYear(year int) bool {
	// Write some code here to pass the test suite.
	if year%400 == 0 {
		return true
	}

	if year%4 == 0 && year%100 != 0 {
		return true
	}

	return false
}
