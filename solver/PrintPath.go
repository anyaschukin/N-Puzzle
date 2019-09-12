package solver

import (
	"fmt"
	g "n-puzzle/golib"
)

func PrintPath(state *State, child []int, size int) {
	var tmp *State
	tmp = nil
	for p := state; p != nil; p = p.before {
		p.path = tmp
		tmp = p
	}
	n := 0
	for p := tmp; p != nil; p = p.path {
		if n == 0 {
			fmt.Println("\x1b[4mInitial state\x1b[0m")
		} else {
			fmt.Printf("\x1b[4mState %d\x1b[0m\n", n)
		}
		g.PrintBoard(p.puzzle, size)
		n++
	}
	fmt.Printf("\x1b[4mState %d - Final solved state\x1b[0m\n", n)
	g.PrintBoard(child, size)
	fmt.Printf("Number of moves required to transition from initial to final state: %d\n", n)
}
