package main

import (
	"fmt"
	"log"
	p "n-puzzle/parsing"
	s "n-puzzle/solver"
	"os"
	"time"
)

func main() {

	start := time.Now()

	size, iterations := p.CheckFlags()
	var Puzzle []int
	//if size == 0 && solve == false && iterations == 0 {
	if size == 0 && iterations == 0 {
		Puzzle, size = p.ReadBoardFromFile(Puzzle, size)
	} else {
		Puzzle = p.GenerateRandomBoard(size)
	}
	s.Solver(Puzzle, size, iterations)

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)

	fmt.Println("\n You've reached the end of main()")

	os.Exit(1)

	//fmt.Println("no file or random board")
}
