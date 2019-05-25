package solver

import (
	"fmt"
)

func swapEmpty(puzzle, size) {
	idx := puzzle[0]

}

func makeGoal(size int) []int {
	totalSize := size * size
	puzzle := make([]int, totalSize)
	cur := 1
	for i := range puzzle {
		if i == (size*size)-1 {
			puzzle[i] = 0
			break
		}
		puzzle[i] = cur
		cur++
	}
	return (puzzle)
}

// if I comment this
func MakePuzzle(size int, solve bool, iterations int) {
	p := makeGoal(size)
	fmt.Print(p)
}
