package main

import (
	"fmt"
	p "n-puzzle/parsing"
	s "n-puzzle/solver"
	"os"
)

func main() {

	// start := time.Now()

	size := p.CheckFlags()
	var Puzzle []int
	if size == 0 {
		Puzzle, size = p.ReadBoardFromFile(Puzzle, size)
	} else {
		Puzzle = p.GenerateRandomBoard(size)
	}
	s.Solver(Puzzle, size)

	// elapsed := time.Since(start)
	// log.Printf("Binomial took %s", elapsed)

	fmt.Println("\n You've finished n-puzzle!")

	os.Exit(1)

	//fmt.Println("no file or random board")
}
