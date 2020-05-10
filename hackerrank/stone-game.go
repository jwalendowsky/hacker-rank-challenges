package hackerrank

import (
	"fmt"
)

// StoneGame is the implementation of the hacker rank challenge of the Stone Games.
// https://www.hackerrank.com/challenges/stonegame/problem
func StoneGame() {
	var numberOfPiles int
	var pile int64

	fmt.Scanf("%d\n", &numberOfPiles)
	piles := make([]int64, numberOfPiles)

	for i := 0; i < numberOfPiles; i++ {
		fmt.Scanf("%d", &pile)
		piles[i] = pile
		fmt.Printf("%064b\n", pile)
	}
}
