package solver

import (
	"fmt"
	g "n-puzzle/golib"
)

// if I comment this
func Solver(Puzzle []int, size int, solve bool, iterations int) {

	Solution := MakeGoal(size)
	g.PrintBoard(Solution, size)

	solveable := IsSolvable(Solution, Puzzle, size)
	fmt.Printf("\nsolve it? %v\n", solveable)

	CreateQueue()

	// apply algo:
	// 	heuristic flags()
	// 	switch case algo1 algo2 algo3
	// 		build tree
	// 		move pieces, generate board per MovePieces
	// 		traverse tree

	// create tree
	// MovePieces(Puzzle)

	goal := g.CheckSliceEquality(Puzzle, Solution)
	fmt.Printf("goal? %v\n", goal)
}

//for i := 1; i <= iterations; i++ {
//	swapEmpty(p, size)
//}
