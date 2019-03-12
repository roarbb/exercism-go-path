// Package leap contains method relevant to leap year detection
package leap

// IsLeapYear determines if supplied year is leap year or not.
func IsLeapYear(year int) bool {

	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return true
	}

	return false
}
