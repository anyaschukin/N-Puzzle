package solver

import (
	g "n-puzzle/golib"
)

func countInversions(puzzle []int) int {
	inversions := 0
	for i := 0; i < len(puzzle)-1; i++ {
		for j := i + 1; j < len(puzzle); j++ {
			if puzzle[i] > puzzle[j] && puzzle[i] != 0 && puzzle[j] != 0 {
				inversions++
			}
		}
	}
	return inversions
}

// if puzzle size is odd
func oddSize(solution []int, puzzle []int, size int) bool {
	pInversions := countInversions(puzzle)
	sInversions := countInversions(solution)
	pIdx := g.FindIndexSlice(puzzle, 0)
	sIdx := g.FindIndexSlice(solution, 0)

	if size%2 == 0 {
		pInversions += pIdx / size
		sInversions += sIdx / size
	}

	return ((pInversions % 2) == (sInversions % 2))
}

func evenSize(solution []int, puzzle []int, size int) bool {
	Inversions := countInversions(puzzle)
	zeroIdx := g.FindIndexSlice(puzzle, 0)
	row := (((size*size - 1) - zeroIdx) / size) + 1
	if ((row%2 == 0) && (Inversions%2 != 0)) || (row%2 != 0) && (Inversions%2 == 0) {
		return true
	}
	return false
}

func IsSolvable(solution []int, puzzle []int, size int) bool {
	if size%2 != 0 || size == 6 || size == 8 {
		return oddSize(solution, puzzle, size)
	} else {
		return evenSize(solution, puzzle, size)
	}
}
