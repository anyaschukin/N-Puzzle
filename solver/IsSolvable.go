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
	// return ((row%2 == 0) && (Inversions%2 != 0)) || ((row%2 != 0) && (Inversions%2 == 0))
}

func IsSolvable(solution []int, puzzle []int, size int) bool {
	if size%2 != 0 || size == 6 || size == 8 {
		return oddSize(solution, puzzle, size)
	} else {
		return evenSize(solution, puzzle, size)
	}
}

// HOLD UP... THIS IS FOR A NON-SNAIL BOARD!! MAYBE NOT THE BEST TO FOLLOW
// FROM https://www.cs.bham.ac.uk/~mdr/teaching/modules04/java2/TilesSolvability.html
// The formula says:
// If the grid width is odd, then the number of inversions in a solvable situation is even.
// If the grid width is even, and the blank is on an even row counting from the bottom (second-last, fourth-last etc), then the number of inversions in a solvable situation is odd.
// If the grid width is even, and the blank is on an odd row counting from the bottom (last, third-last, fifth-last etc) then the number of inversions in a solvable situation is even.
// That gives us this formula for determining solvability:
//
// ( (grid width odd) && (#inversions even) )  ||  ( (grid width even) && ((blank on odd row from bottom) == (#inversions even)) )

// from Damien
// If N is even, position of 0 from bottom is needed
// else {
// const row = find(this.puzzle, { value: 0 });
// const snailRow = find(this.snail, { value: 0 });
// const numberOfRows = Math.abs(snailRow.y - row.y) + Math.abs(snailRow.x - row.x);
// this.solvable = (numberOfRows % 2 !== numberOfPermutations % 2);
// }
