package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// returns a random boolean
func rand1() bool {
	return rand.Float32() < 0.5
}

//func (u *UniqueRand) Int() int {
//	for {
//		i := rand.Int()
//		if !u.generated[i] {
//			u.generated[i] = true
//			return i
//		}
//	}
//}

func makeRangeNum(min, max int) []int {
	set := make([]int, max-min+1)
	for i := range set {
		set[i] = min + i
	}
	return set
}

func generateRandomBoard(size int, solve bool, iterations int) {
	// generate a shuffled set of numbers
	maxNb := size * size
	numbers := makeRangeNum(1, maxNb)
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})

	// generate board
	sizey := size
	sizex := size
	puzzle := make([][]int, sizey)
	for i := range puzzle {
		puzzle[i] = make([]int, sizex)

	}

	// fill board
	index := 0
	for i := 0; i < sizex; i++ {
		for j := 0; j < sizey; j++ {
			puzzle[i][j] = numbers[index]
			index++
		}
	}

	fmt.Println(puzzle)

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

	//	_ = solve

	size := *sizePtr
	iterations := *iterationsPtr

	generateRandomBoard(size, solve, iterations)
}
