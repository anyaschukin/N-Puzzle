package solver

import (
	"fmt"
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
	length := width * width
	// if zero is in top row
	if i < length-width {
		down = 1
	}
	// if zero is not in top row
	if i >= width && i < length {
		up = 1
	}
	// if zero is in right column
	maxMod := (length - 1) % width
	// if zero is in left column
	minMod := 0
	if i%width != minMod {
		left = 1
	}
	if i%width != maxMod {
		right = 1
	}
	return up, down, left, right
}

func MovePieces(puzzle []int, size int) {

	empty := g.FindIndexSlice(puzzle, 0)
	up, down, left, right := checkMoves(size, empty)
	fmt.Printf("up %d, down %d, left %d, right %d\n", up, down, left, right)

	moves := up + down + left + right
	fmt.Printf("\nthere are %d possible moves\n", moves)
	neighbors := map[string][]int{}
	g.PrintBoard(puzzle, size)
	fmt.Printf("----\n")
	for moves > 0 {
		key := fmt.Sprintf("neighbor%d", moves)
		neighbor := make([]int, len(puzzle))
		copy(neighbor, puzzle)
		neighbors[key] = neighbor
		// move pieces
		// use switch for up, down, left, right?
		if up == 1 {
			fmt.Printf("UP empty = %d, empty-size= %d", empty, empty-size)
			neighbor[empty], neighbor[empty-size] = neighbor[empty-size], neighbor[empty]
			up = 0
		} else if down == 1 {
			fmt.Printf("DOWN empty = %d, empty-size= %d", empty, empty+size)
			neighbor[empty], neighbor[empty+size] = neighbor[empty+size], neighbor[empty]
			down = 0
		} else if left == 1 {
			fmt.Printf("LEFT empty = %d, empty-size= %d", empty, empty-size)
			neighbor[empty], neighbor[empty-1] = neighbor[empty-1], neighbor[empty]
			left = 0
		} else if right == 1 {
			fmt.Printf("RIGHT empty = %d, empty-size= %d", empty, empty-size)
			neighbor[empty], neighbor[empty+1] = neighbor[empty+1], neighbor[empty]
			right = 0
		}
		moves--
		fmt.Printf("\n %s :", key)
		g.PrintBoard(neighbor, size)
		fmt.Println("----")
	}
	//return new

	//	for i := 0; i <= moves; i++ {
	//		new := make([]int, size*size)
	// create new board according to each possible move,
	// to store in decision tree
	//	}
}
