package main

import (
	"fmt"
	p "n-puzzle/parsing"
	s "n-puzzle/solver"
	"os"
)

func main() {

	size, solve, iterations := p.CheckFlags()
	if size == 0 && solve == false && iterations == 0 {
		p.ReadBoardFromFile()
	} else {
		p.GenerateRandomBoard(size, solve, iterations)
	}
	s.MakePuzzle(size, solve, iterations)
	fmt.Println("\n\n\n You've reached the end of main()")
	os.Exit(1)

	//fmt.Println("no file or random board")
}
