package parsing

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	g "n-puzzle/golib"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func GenerateRandomBoard(size int) []int {
	//	fmt.Println("\nGenerating random board...\n")

	// generate a shuffled set of numbers
	maxNb := size * size
	numbers := g.MakeRangeNum(0, maxNb-1)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})
	// generate + fill board
	Puzzle := make([]int, maxNb)
	index := 0
	for i := 0; i < maxNb; i++ {
		Puzzle[i] = numbers[index]
		index++
	}
	//	fmt.Println(Puzzle)
	Puzzle = g.SpiralMatrix(Puzzle, size)
	return Puzzle
}

func ReadBoardFromFile(Puzzle []int, size int) ([]int, int) {

	args := os.Args[1:]
	arg := strings.Join(args, "")
	file, err := ioutil.ReadFile(arg)
	g.Check(err)

	if strings.ContainsAny(string(file), ".") || strings.ContainsAny(string(file), "-") {
		fmt.Println("Error: Board values cannot be negative or floats.")
		os.Exit(1)
	}

	re := regexp.MustCompile("[-+]?[0-9]+") // finds all numbers, including negative
	// re := regexp.MustCompile("[-+]?[0-9]*\\.?[0-9]+") // finds all numbers, including negative and floats
	// re := regexp.MustCompile("[^1-9]") // finds all non-numbers

	numbers := re.FindAllString(string(file), -1)
	fmt.Print("puzzle = %s", numbers)

	// convert []string array to []int slice
	i := -1
	for _, number := range numbers {
		if i++; i == 0 {
			size, err = strconv.Atoi(number)
			g.Check(err)
			continue
		}
		integer, err := strconv.Atoi(number)
		g.Check(err)
		Puzzle = append(Puzzle, integer)
	}

	var filtered int
	Puzzle, filtered = g.FilterDuplicates(Puzzle)
	if ((i - filtered) % size) != 0 {
		fmt.Print("\n Not enough pieces to fill board!\n")
		os.Exit(1)
	}
	Puzzle = g.SpiralMatrix(Puzzle, size)
	g.PrintBoard(Puzzle, size)
	return Puzzle, size
}

func CheckFlags() (int, int) {
	sizePtr := flag.Int("size", 1, "Size of the puzzle's side. Must be >3.")
	//solveablePtr := flag.Bool("s", false, "Forces generation of a solvable puzzle. Overrides -u.")
	//unsolveablePtr := flag.Bool("u", false, "Forces generation of an unsolvable puzzle")
	iterationsPtr := flag.Int("iterations", 100, "Number of passes")

	flag.Parse()
	args := flag.Args()

	arg := strings.Join(args, "")
	file := strings.Contains(arg, ".txt")

	if len(args) == 1 && file {
		return 0, 0
	}

	if len(args) > 1 && file {
		fmt.Println("Error: must input one file OR flags as argument.")
		os.Exit(1)
	}

	if *sizePtr < 3 {
		flag.PrintDefaults() // replace with Print Usage
		os.Exit(1)
	}

	if sizePtr == nil {
		fmt.Println("Error: please give a board size.")
		os.Exit(1)
	}

	//if *solveablePtr && *unsolveablePtr {
	//	fmt.Println("Can't be both solvable AND unsolvable, dummy!")
	//	os.Exit(1)
	//}

	if *iterationsPtr < 1 {
		fmt.Println("Can't solve a puzzle in less than 1 iteration!")
		os.Exit(1)
	}

	//var solve bool
	//if !(*solveablePtr) && !(*unsolveablePtr) {
	//	rand.Seed(time.Now().UnixNano())
	//	solve = g.RandomBool()
	//} else if *solveablePtr {
	//	solve = true
	//} else if *unsolveablePtr {
	//	solve = false
	//}

	size := *sizePtr
	iterations := *iterationsPtr

	return size, iterations
}
