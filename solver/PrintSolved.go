package solver

import (
	"fmt"
	"time"
)

func PrintSolved(elapsed time.Duration, problem Problem, state *State, child []int, size int) {
	fmt.Println("This puzzle has been solved!\n")
	PrintPath(state, child, size)
	fmt.Printf("Size Complexity: %d\n", problem.sizeComplexity)
	fmt.Printf("Time Complexity: %d\n", problem.timeComplexity)
	fmt.Printf("Binomial took %s", elapsed)
	fmt.Println("\nYou've finished n-puzzle!")
}
