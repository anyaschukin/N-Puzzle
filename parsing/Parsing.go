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
	return Puzzle
}

func ReadBoardFromFile(Puzzle []int, size int, flags int) ([]int, int) {

	args := os.Args[flags+1:]
	arg := strings.Join(args, "")
	file, err := ioutil.ReadFile(arg)
	g.Check(err, "Error reading file.\nUsage: $> ./n-puzzle -h={heuristic} file.txt")

	if strings.ContainsAny(string(file), ".") || strings.ContainsAny(string(file), "-") {
		fmt.Println("Error: Board values cannot be negative or floats.")
		os.Exit(1)
	}

	re := regexp.MustCompile("[-+]?[0-9]+") // finds all numbers, including negative
	// re := regexp.MustCompile("[-+]?[0-9]*\\.?[0-9]+") // finds all numbers, including negative and floats
	// re := regexp.MustCompile("[^1-9]") // finds all non-numbers

	numbers := re.FindAllString(string(file), -1)

	// convert []string array to []int slice
	i := -1
	for _, number := range numbers {
		if i++; i == 0 {
			size, err = strconv.Atoi(number)
			g.Check(err, "Error board contains non-integers")
			continue
		}
		integer, err := strconv.Atoi(number)
		g.Check(err, "Error board contains non-integers")
		Puzzle = append(Puzzle, integer)
	}

	var filtered int
	Puzzle, filtered = g.FilterDuplicates(Puzzle)
	if ((i - filtered) % size) != 0 {
		fmt.Print("\n Not enough pieces to fill board!\n")
		os.Exit(1)
	}
	return Puzzle, size
}

func CheckFlags() (int, string, int) {
	sizePtr := flag.Int("s", 1, "size of the puzzle's side must be > 3.")
	heuristicPtr := flag.String("h", "manhattan", "Heuristic options include: manhattan, hamming, euclidian, nilsson, and outRowCol.")

	flag.Parse()
	args := flag.Args()
	flags := flag.NFlag()

	arg := strings.Join(args, "")
	file := strings.Contains(arg, ".txt")

	heuristic := *heuristicPtr
	switch heuristic {
	case "manhattan":
	case "hamming":
	case "euclidean":
	case "nilsson":
	case "outRowCol":
	default:
		heuristic = "manhattan"
	}

	if file {
		return 0, heuristic, flags
	}

	if *sizePtr < 3 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if sizePtr == nil {
		fmt.Println("Error: please give a board size.")
		os.Exit(1)
	}

	size := *sizePtr

	return size, heuristic, flags
}
