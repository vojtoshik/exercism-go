package scrabble

import (
	"strings"
)

var scores = []struct {
	letters string
	score   int
}{
	{"AEIOULNRST", 1},
	{"DG", 2},
	{"BCMP", 3},
	{"FHVWY", 4},
	{"K", 5},
	{"JX", 8},
	{"QZ", 10},
}

func Score(s string) int {

	scoreMap := createScoreMap()
	var result = 0

	for _, letter := range strings.ToUpper(s) {
		result += scoreMap[letter]
	}

	return result
}

func createScoreMap() map[int32]int {
	letterToScore := make(map[int32]int)

	for _, score := range scores {
		for _, letter := range score.letters {

			letterToScore[letter] = score.score
		}
	}

	return letterToScore
}
