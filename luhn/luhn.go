package luhn

import (
	"regexp"
	"strings"
)

const AsciiCodeZero = 48

func Valid(ccNumber string) bool {

	ccNumber = strings.Replace(ccNumber, " ", "", -1)
	re := regexp.MustCompile("^[0-9]{2,}$")

	if !re.MatchString(ccNumber) {
		return false
	}

	sum := 0

	for i:=1; i <= len(ccNumber); i++ {

		digit := int(ccNumber[len(ccNumber) - i]) - AsciiCodeZero

		if i % 2 == 1 {
			sum += digit
			continue
		}

		v := 2 * digit

		if v >= 10 {
			v -= 9
		}

		sum += v
	}

	return sum % 10 == 0
}