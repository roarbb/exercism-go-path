// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import (
	"math"
)

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = "NaT" // not a triangle
	Equ Kind = "Equ" // equilateral
	Iso Kind = "Iso" // isosceles
	Sca Kind = "Sca" // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {

	var k Kind

	anyIsNegative := func(a, b, c float64) bool {
		return a < 0 || b < 0 || c < 0
	}

	isTriangle := func(a, b, c float64) bool {
		if a+b < c || a+c < b || c+b < a {
			return false
		}

		return true
	}

	allSidesZero := func(a, b, c float64) bool {
		return a+b+c == 0
	}

	allSidesAreNumbers := func(a, b, c float64) bool {
		return !math.IsNaN(a) && !math.IsNaN(b) && !math.IsNaN(c) && !math.IsInf(a, 0) && !math.IsInf(b, 0) && !math.IsInf(c, 0)
	}

	switch {
	case a == 0 && b == 0 && c == 0 || anyIsNegative(a, b, c) || !isTriangle(a, b, c) || allSidesZero(a, b, c) || !allSidesAreNumbers(a, b, c):
		k = NaT
	case a == b && a == c:
		k = Equ
	case (a == b || a == c || b == c) && (a != b || a != c || b != c):
		k = Iso
	case a != b && b != c && c != a:
		k = Sca
	}

	return k
}
