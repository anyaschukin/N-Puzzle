package solver

import (
	"fmt"
	g "n-puzzle/golib"
)

func countInversions(puzzle []int) int {
	inversions := 0
	for i := 0; i < len(puzzle); i++ {
		for j := i + 1; j < len(puzzle); j++ {
			if j < i {
				inversions++
			}
		}
	}
	return inversions
}

// need to test this
func IsSolvable(solution []int, puzzle []int, size int) bool {
	pInversions := countInversions(puzzle)
	sInversions := countInversions(solution)
	pIdx := g.FindIndexSlice(puzzle, 0)
	sIdx := g.FindIndexSlice(solution, 0)

	if size%2 == 0 {
		pInversions += pIdx / size
		sInversions += sIdx / size
	}

	if (pInversions % 2) != (sInversions % 2) {
		fmt.Println("This board is unsolvable :( \n")
		return false
	} else {
		fmt.Println("This board is solvable :) \n")
		return true
	}
}
