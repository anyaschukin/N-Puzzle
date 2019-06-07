package solver

import (
	g "n-puzzle/golib"
)

func swapEmpty(puzzle []int, size int) {
	//fmt.Print(puzzle)
	idx := g.FindIndexSlice(puzzle, 0)
	//fmt.Print(idx)
	//fmt.Println("\n")
	var Poss []int
	if idx%size > 0 {
		Poss = append(Poss, idx-1)
	}
	if idx%size < size-1 {
		Poss = append(Poss, idx+1)
	}
	if idx/size > 0 {
		Poss = append(Poss, idx-size)
	}
	if idx/size < size-1 {
		Poss = append(Poss, idx+size)
	}
	//fmt.Print(Poss)
	swi := g.RandomChoice(Poss)
	//fmt.Print(swi)
	puzzle[idx] = puzzle[swi]
	puzzle[swi] = 0
}

// if I comment this
func MakePuzzle(size int, solve bool, iterations int) {
	p := MakeGoal(size)
	g.PrintBoard(p, size)

	//for i := 1; i <= iterations; i++ {
	//	swapEmpty(p, size)
	//}
}
