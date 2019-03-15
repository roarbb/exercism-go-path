// Package hamming helps with counting the DNA differences
package hamming

import (
	"errors"
)

// Distance returns hamming distance between 2 DNAs
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return -1, errors.New("Error: Different length of strings")
	}

	var distance int

	for index := 0; index < len(a); index++ {
		if a[index] != b[index] {
			distance++
		}
	}

	return distance, nil
}
