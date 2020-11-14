package letter

import "math"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// ConcurrentFrequency calculates the frequency of runes concurrently for
// each text passed as parameter
func ConcurrentFrequency(texts []string) FreqMap {
	c := make(chan FreqMap, calculateChannelSize(len(texts)))
	for _, text := range texts {
		go func(s string) {
			c <- Frequency(s)
		}(text)
	}

	subtotals := make(FreqMap)

	for range texts {

		var freqMap = <-c

		for k, v := range freqMap {
			subtotals[k] += v
		}
	}

	return subtotals
}

func calculateChannelSize(tasksNumber int) int {
	tn := float64(tasksNumber)
	return int(math.Sqrt(tn) + tn/3 + 1)
}

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}
