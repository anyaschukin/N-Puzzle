package main

import (
	p "n-puzzle/parsing"
	s "n-puzzle/solver"
)

func main() {

	size, heuristic, flags := p.CheckFlags()
	var Puzzle []int
	if size == 0 {
		Puzzle, size = p.ReadBoardFromFile(Puzzle, size, flags)
	} else {
		Puzzle = p.GenerateRandomBoard(size)
	}
	s.Solver(Puzzle, size, heuristic)
}
