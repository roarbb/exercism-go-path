// Package triangle contains methods to work with triangle shape
package triangle

import (
	"math"
)

// Kind holds the type of triangle
type Kind string

const (
	// NaT defines not a triangle
	NaT Kind = "NaT"
	// Equ defines equilateral triangle
	Equ Kind = "Equ"
	// Iso defines isosceles triangle
	Iso Kind = "Iso"
	// Sca defines scalene triangle
	Sca Kind = "Sca"
)

// KindFromSides returns the kind of triangle based on triangle sides measurements
func KindFromSides(a, b, c float64) Kind {

	var k Kind

	isNegative := func(a, b, c float64) bool {
		return a < 0 || b < 0 || c < 0
	}

	isTriangle := func(a, b, c float64) bool {
		return a+b >= c && a+c >= b && c+b >= a
	}

	allSidesZero := func(a, b, c float64) bool {
		return a+b+c == 0
	}

	allSidesAreNumbers := func(a, b, c float64) bool {
		return !math.IsNaN(a) && !math.IsNaN(b) && !math.IsNaN(c) && !math.IsInf(a, 0) && !math.IsInf(b, 0) && !math.IsInf(c, 0)
	}

	allSidesEqual := func(a, b, c float64) bool {
		return a == b && a == c
	}

	twoSidesEqual := func(a, b, c float64) bool {
		return (a == b || a == c || b == c) && (a != b || a != c || b != c)
	}

	allSidesDifferent := func(a, b, c float64) bool {
		return a != b && b != c && c != a
	}

	switch {
	case allSidesZero(a, b, c) || isNegative(a, b, c) || !isTriangle(a, b, c) || !allSidesAreNumbers(a, b, c):
		k = NaT
	case allSidesEqual(a, b, c):
		k = Equ
	case twoSidesEqual(a, b, c):
		k = Iso
	case allSidesDifferent(a, b, c):
		k = Sca
	}

	return k
}
