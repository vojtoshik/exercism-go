package acronym

import (
	"regexp"
	"strings"
)

func Abbreviate(s string) string {

	re := regexp.MustCompile("[\\s-_]+")
	words := re.Split(s, -1)

	var result = ""

	for _, word := range words {
		result += strings.ToUpper(word[:1])
	}

	return result
}
