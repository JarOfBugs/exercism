// Leap contains function about leap years
package leap

// IsLeapYear Check if the given year is a leap year
// returns true if the given year is leap, false otherwise
func IsLeapYear(year int) bool {
	return year % 4 == 0 && year % 100 != 0 || year % 100 == 0 && year % 400 == 0
}
