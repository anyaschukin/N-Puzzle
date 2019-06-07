package solver

import (
	"math/rand"
	g "n-puzzle/golib"
	"time"
)

func randomChoice(Poss []int) int {
	rand.Seed(time.Now().Unix())
	swi := rand.Intn(len(Poss))
	return swi
}

func findIndexSlice(slice []int, value int) int {
	for p, v := range slice {
		if value == v {
			return p
		}
	}
	return -1
}

func swapEmpty(puzzle []int, size int) {
	//fmt.Print(puzzle)
	idx := findIndexSlice(puzzle, 0)
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
	swi := randomChoice(Poss)
	//fmt.Print(swi)
	puzzle[idx] = puzzle[swi]
	puzzle[swi] = 0
}

// if I comment this
func MakePuzzle(size int, solve bool, iterations int) {
	p := MakeGoal(size)
	//for i := 1; i <= iterations; i++ {
	//	swapEmpty(p, size)
	//}
	g.PrintBoard(p, size)
}
