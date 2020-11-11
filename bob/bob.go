package bob

import (
	"strings"
	"unicode"
)

func Hey(remark string) string {

	isPhrase, isYelling := phraseAnalysis(remark)
	isQuestion := isQuestion(remark)

	if isYelling && isQuestion {
		return "Calm down, I know what I'm doing!"
	}

	if isQuestion {
		return "Sure."
	}

	if isYelling {
		return "Whoa, chill out!"
	}

	if !isPhrase {
		return "Fine. Be that way!"
	}

	return "Whatever."
}

func isQuestion(s string) bool {
	s = strings.Trim(s, " ")
	return len(s) > 0 && s[len(s)-1] == '?'
}

func phraseAnalysis(s string) (bool, bool) {

	containsNumbers := false
	containsLetters := false
	containsLowerCaseLetters := false

	r := []rune(s)

	for i := 0; i < len(r); i++ {

		l := r[i]

		if unicode.IsNumber(l) {
			containsNumbers = true
			continue
		}

		if unicode.IsLetter(l) {
			containsLetters = true

			if unicode.IsLower(l) {
				containsLowerCaseLetters = true
			}
		}
	}

	return containsLetters || containsNumbers, containsLetters && !containsLowerCaseLetters
}