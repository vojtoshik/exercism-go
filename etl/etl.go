package etl

import "strings"

func Transform (input map[int][]string) (r map[string]int) {

	r = make(map[string]int)

	for p, lc := range input {
		for _, l := range lc {
			r[strings.ToLower(l)] = p
		}
	}

	return r
}