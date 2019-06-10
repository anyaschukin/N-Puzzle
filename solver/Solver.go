package solver

import (
	"fmt"
	g "n-puzzle/golib"
)

//func swapEmpty(puzzle []int, size int) {
//	//fmt.Print(puzzle)e
//	idx := g.FindIndexSlice(puzzle, 0)
//	//fmt.Print(idx)
//	//fmt.Println("\n")
//	var Poss []int
//	if idx%size > 0 {
//		Poss = append(Poss, idx-1)
//	}
//	if idx%size < size-1 {
//		Poss = append(Poss, idx+1)
//	}
//	if idx/size > 0 {
//		Poss = append(Poss, idx-size)
//	}
//	if idx/size < size-1 {
//		Poss = append(Poss, idx+size)
//	}
//	//fmt.Print(Poss)
//	swi := g.RandomChoice(Poss)
//	//fmt.Print(swi)
//	puzzle[idx] = puzzle[swi]
//	puzzle[swi] = 0
//}

// if size is odd
//	if size % 2 != 0 {
//		if pInversions % 2 == 0 {

// if I comment this
func MakePuzzle(Puzzle []int, size int, solve bool, iterations int) {

	p := MakeGoal(size)
	g.PrintBoard(p, size)

	solveable := IsSolvable(p, Puzzle, size)

	fmt.Printf("\nsolve it? %v\n", solveable)
	//for i := 1; i <= iterations; i++ {
	//	swapEmpty(p, size)
	//}
}
