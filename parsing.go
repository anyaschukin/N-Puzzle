package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func rand1() bool {
	return rand.Float32() < 0.5
}

func main() {
	//	argsWithProg := os.Args
	//	argsWithoutProg := os.Args[1:]

	//	arg = os.Args[3]

	sizePtr := flag.Int("size", 1, "Size of the puzzle's side. Must be >3.")
	solveablePtr := flag.Bool("s", false, "Forces generation of a solvable puzzle. Overrides -u.")
	unsolveablePtr := flag.Bool("u", false, "Forces generation of an unsolvable puzzle")
	iterationsPtr := flag.Int("iterations", 10000, "Number of passes")

	flag.Parse()

	//	fmt.Println("size:", *sizePtr)

	if *sizePtr < 3 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *solveablePtr && *unsolveablePtr {
		fmt.Println("Can't be both solvable AND unsolvable, dummy!")
		os.Exit(1)
	}

	if *iterationsPtr < 1 {
		fmt.Println("Can't solve a puzzle in less than 1 iteration!")
		os.Exit(1)
	}

	var solve bool

	if !(*solveablePtr) && !(*unsolveablePtr) {
		rand.Seed(time.Now().UnixNano())
		solve = rand1()
	} else if *solveablePtr {
		solve = true
	} else if *unsolveablePtr {
		solve = false
	}

	//size := *sizePtr

	//	puzzle := make_puzzle(size, )
}
