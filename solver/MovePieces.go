package solver

import (
	g "n-puzzle/golib"
)

// look at current board for potential moves
// count n potential moves
// check if move is valid
// generate n boards for n moves
// ... how to check if board is duplicate?

// Checks which directions the zero tile can move in
func checkMoves(width int, i int) (int, int, int, int) {
	up, down, left, right := 0, 0, 0, 0
	// if zero is in top row
	if i <= width {
		down = 1
	}
	// if zero is in bottom row
	length := width * width
	if i >= length-width && i <= length {
		up = 1
	}
	// if zero is in left column
	if i%width == 1 {
		right = 1
	}
	// if zero is in right column
	if i%width == 0 {
		left = 1
	}
	return up, down, left, right
}

func MovePieces(puzzle []int, size int) []int {

	empty := g.FindIndexSlice(puzzle, 0)
	up, down, left, right := checkMoves(size, empty)

	//return new
}

// for i := n moves {
//	new := make([]int, size*size)
//}
