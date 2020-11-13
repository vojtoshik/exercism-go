package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// ConcurrentFrequency calculates the frequency of runes concurrently for
// each text passed as parameter
func ConcurrentFrequency(texts []string) FreqMap {
	c := make(chan FreqMap)
	for _, text := range texts {
		go goroutineFrequency(text, c)
	}

	subtotals := make(FreqMap)

	for i := 0; i < len(texts); i++ {

		var freqMap = <-c

		for k, v := range freqMap {
			subtotals[k] += v
		}
	}

	return subtotals
}

func goroutineFrequency(s string, c chan<- FreqMap) {
	c <- Frequency(s)
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
