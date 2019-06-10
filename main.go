package main

import (
	"fmt"
	p "n-puzzle/parsing"
	s "n-puzzle/solver"
	"os"
)

func main() {

	size, solve, iterations := p.CheckFlags()
	var Puzzle []int
	if size == 0 && solve == false && iterations == 0 {
		Puzzle = p.ReadBoardFromFile()
	} else {
		Puzzle = p.GenerateRandomBoard(size, solve, iterations)
		//s.MovePieces(puzzle, size)
	}
	s.MakePuzzle(Puzzle, size, solve, iterations)
	fmt.Println("\n\n\n You've reached the end of main()")
	os.Exit(1)

	//fmt.Println("no file or random board")
}
