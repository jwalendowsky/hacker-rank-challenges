package hackerrank

import (
	"bufio"
	"fmt"
	"os"
)

// MakingCandies solves the Hacker rank problem of https://www.hackerrank.com/challenges/making-candies/problem.
func MakingCandies() {
	var machines, workers, price, quantityToProduce int
	reader := bufio.NewReader(os.Stdin)
	var numberOfPasses int = 0
	var quantityProduced int = 0

	fmt.Fscanf(reader, "%d %d %d %d\n", &machines, &workers, &price, &quantityToProduce)

	for quantityProduced < quantityToProduce {
		// production := machines * workers
		// maximumToBuy := production / price
		numberOfPasses++
	}

	fmt.Println(numberOfPasses)
}
