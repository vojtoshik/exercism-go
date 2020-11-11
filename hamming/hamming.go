package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("strings are not of the same length")
	}

	result := 0

	for i,_ := range a {
		if a[i:i+1] != b[i:i+1] {
			result++
		}
	}

	return result, nil
}
