package raindrops

import "strconv"

var factors = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

func Convert(n int) string  {
	result := ""

	for factor,sound := range factors {
		if n % factor == 0 {
			result += sound
		}
	}

	if result == "" {
		result = strconv.Itoa(n)
	}

	return result
}