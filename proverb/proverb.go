// Package proverb generates proverb
// Example:
// For want of a nail the shoe was lost.
// For want of a shoe the horse was lost.
// For want of a horse the rider was lost.
// For want of a rider the message was lost.
// For want of a message the battle was lost.
// For want of a battle the kingdom was lost.
// And all for the want of a nail.
package proverb

import "fmt"

// Proverb generates the relevant proverb based on a list of inputs
func Proverb(rhyme []string) (proverb []string) {
	if len(rhyme) == 0 {
		return
	}

	for index := 1; index < len(rhyme); index++ {
		proverb = append(proverb, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[index-1], rhyme[index]))
	}

	proverb = append(proverb, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	return
}
