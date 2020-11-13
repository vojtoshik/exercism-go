package isogram

import (
	"regexp"
	"strings"
)

func IsIsogram(s string) bool {

	re, _ := regexp.Compile("\\W")
	s = re.ReplaceAllString(strings.ToLower(s), "")

	for i, _ := range s {
		if strings.Contains(s[i+1:], s[i:i+1]) {
			return false
		}
	}

	return true
}
