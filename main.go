package main

import (
	"fmt"
	p "n-puzzle/parsing"
	"os"
)

func main() {

	size, solve, iterations := p.CheckFlags()
	if size == 0 && solve == false && iterations == 0 {
		p.ReadBoardFromFile()
	} else {
		p.GenerateRandomBoard(size, solve, iterations)
	}
	fmt.Println("\n\n\n You've reached the end of main()\n")
	os.Exit(1)

	//fmt.Println("no file or random board")
}
