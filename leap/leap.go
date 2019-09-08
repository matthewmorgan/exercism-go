// Package leap has a simple method for determining leap years.
package leap

// IsLeapYear accepts a year as an integer, and returns true if it is a leap year, false otherwise.
func IsLeapYear(year int) bool {
	switch {
		case year%4 != 0:
			return false
		case year%100 == 0 && year%400 != 0:
			return false
	}
	return true
}
