// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
    NaT = iota
    Equ
    Iso
    Sca
)

// KindFromSides returns type of triangle that sides of length a, b and c compose
func KindFromSides(a, b, c float64) Kind {

	if !isValidValue(a) || !isValidValue(b) || !isValidValue(c) || a + b < c || b + c < a || c + a < b {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || b == c || c == a {
		return Iso
	}

	return Sca
}

func isValidValue(v float64) bool {
	return v > 0 && !math.IsInf(v, 1)
}