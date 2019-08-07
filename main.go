package main

import (
	"fmt"
	p "n-puzzle/parsing"
	s "n-puzzle/solver"
	"os"
)

func main() {

	size, iterations := p.CheckFlags()
	var Puzzle []int
	//if size == 0 && solve == false && iterations == 0 {
	if size == 0 && iterations == 0 {
		Puzzle, size = p.ReadBoardFromFile(Puzzle, size)
	} else {
		Puzzle = p.GenerateRandomBoard(size)
		//s.MovePieces(puzzle, size)
	}
	s.Solver(Puzzle, size, iterations)
	fmt.Println("\n You've reached the end of main()")
	os.Exit(1)

	//fmt.Println("no file or random board")
}
