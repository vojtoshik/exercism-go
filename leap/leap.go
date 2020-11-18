package leap

// IsLeapYear returns if the year is leap year
func IsLeapYear(year int) bool {

	return year%4 == 0 && (year%100 != 0 || year%100 == 0 && year%400 == 0)
}
