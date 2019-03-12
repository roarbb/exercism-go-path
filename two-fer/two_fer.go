// Package twofer provides methods that helps you with sharing.
package twofer

import "fmt"

// ShareWith returns 'two for one' sentence.
// If the name is not provided it will return default string, otherwise it'll use the name.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %v, one for me.", name)
}
