package twelve

import (
	"fmt"
	"strings"
)

var presents = []string{
	"twelve Drummers Drumming",
	"eleven Pipers Piping",
	"ten Lords-a-Leaping",
	"nine Ladies Dancing",
	"eight Maids-a-Milking",
	"seven Swans-a-Swimming",
	"six Geese-a-Laying",
	"five Gold Rings",
	"four Calling Birds",
	"three French Hens",
	"two Turtle Doves",
	"a Partridge in a Pear Tree",
}

var ordinals = []string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

// Song returns lyrics of the song "Twelve Days"
func Song() (s string) {

	for i := 1; i <= 12; i++ {

		if s != "" {
			s += "\n"
		}

		s += Verse(i)
	}

	return s
}

// Verse returns line of the song about a particular day of the Christmas
func Verse(day int) string {

	day = day - 1
	r := fmt.Sprintf("On the %s day of Christmas my true love gave to me: ",
		ordinals[day])

	if day == 0 {
		r += presents[len(presents)-1]
	} else {
		r += strings.Join(presents[len(presents)-day-1:len(presents)-1], ", ") +
			", and " + presents[len(presents)-1]
	}

	return r + "."
}
