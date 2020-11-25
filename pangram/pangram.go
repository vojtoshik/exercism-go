package pangram

import (
	"regexp"
	"strings"
)

const lettersInLatinAlphabet = 26

// IsPangram returns true if the input is a pangram
func IsPangram(in string) bool {
	in = strings.ToLower(in)

	in = regexp.MustCompile("[^a-z]").ReplaceAllString(in, "")

	if in == "" {
		return false
	}

	h := make(map[int]bool)
	uniqueLetters := 0

	for _, r := range in {
		if _,ok := h[int(r)]; !ok {
			h[int(r)] = true
			uniqueLetters++
		}
	}

	return uniqueLetters == lettersInLatinAlphabet
}
