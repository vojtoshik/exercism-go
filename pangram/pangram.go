package pangram

import (
	"regexp"
	"strings"
)

// IsPangram returns true if the input is a pangram
func IsPangram(in string) bool {
	in = strings.ToLower(in)

	in = regexp.MustCompile("[^a-z]").ReplaceAllString(in, "")

	if in == "" {
		return false
	}

	h := make(map[int]bool)

	for _, r := range in {
		h[int(r)] = true
	}

	if len(h) != 26 {
		return false
	}

	for _, v := range h {
		if !v {
			return false
		}
	}

	return true
}
