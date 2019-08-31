package solver

import (
	"container/heap"
	"fmt"
	g "n-puzzle/golib"
	"os"
	"reflect"
	// "time"
	"github.com/AndreasBriese/bbloom"
)

type Problem struct {
	start    []int
	goal     []int
	solvable bool
	//heuristic      string
	//searchAlgo     string
	solutionPath   map[int][]int
	sizeComplexity int
	timeComplexity int
	solutionFound  bool
}

func newProblem(Puzzle []int, size int) Problem {
	problem := Problem{}
	problem.start = Puzzle
	problem.goal = MakeGoal(size)
	problem.solvable = IsSolvable(problem.goal, Puzzle, size)
	//problem.heuristic = "MANHATTAN"
	//problem.searchAlgo = "A_STAR"
	problem.sizeComplexity = 0
	problem.timeComplexity = 1
	return problem
}

type State struct {
	index     int
	priority  int
	depth     int
	heuristic int
	puzzle    []int
}

func newState(Puzzle []int, size int, priority int, depth int, heuristic int) *State {
	state := &State{}
	state.index = 0
	state.priority = priority
	state.depth = depth
	state.heuristic = heuristic
	state.puzzle = Puzzle
	return state
}

func Solver(Puzzle []int, size int, iterations int) {
	problem := newProblem(Puzzle, size)

	// g.PrintBoard(Puzzle, size)

	if IsSolvable(problem.goal, Puzzle, size) == false {
		fmt.Println("This puzzle in unsolvable.")
		os.Exit(1)
	}
	closedSet := bbloom.New(float64(1<<16), float64(0.01))
	state := newState(Puzzle, size, 10000, 0, 0)
	openQueue := CreateQueue(*state)

	for counter := 0; counter < 6000000; counter++ {
		
		if len(openQueue) == 0 {
			fmt.Println("This priorityQueue is empty.")
			g.PrintBoard(state.puzzle, size)
			os.Exit(1)
		}
		state = heap.Pop(&openQueue).(*State)
		closedSet.AddIfNotHas([]byte(g.PuzzleToString(state.puzzle, ",")))

		// fmt.Println(" ----------- ")
		// fmt.Println("\n NEW STATE")
		// fmt.Printf("\n priority = %d, heuristic = %d, depth = %d", state.priority, state.heuristic, state.depth)
		// g.PrintBoard(state.puzzle, size)
		// fmt.Println(" ----------- ")

		if reflect.DeepEqual(problem.goal, state.puzzle) {
			fmt.Println("This puzzle has been solved!\n")
			g.PrintBoard(state.puzzle, size)
			// REBUILD PATH TO START
			os.Exit(1)
		}

		// time.Sleep(1000 * time.Millisecond)

		children := CreateNeighbors(state.puzzle, size)
		// fmt.Println("--- children ---")
		// fmt.Print(children)

		// fmt.Println("\n CHILDREN \n")
		for _, child := range children {
			// g.PrintBoard(child, size)

			if closedSet.Has([]byte(g.PuzzleToString(child, ","))) {
				problem.sizeComplexity++
				continue
			}
			problem.timeComplexity++
			// s := newState(child, size, priority, state.depth+1, heuristic)

			heuristic := g.Manhattan(child, problem.goal, size)
			priority := state.depth + 1 + heuristic
			// priority = -priority
			// priority :=  heuristic
			s := newState(child, size, priority, state.depth+1, heuristic)
			// if s.priority > state.priority {
			// 	continue
			// }
			// fmt.Printf("\n priority = %d, heuristic = %d, depth = %d", priority, heuristic, state.depth + 1)
			// g.PrintBoard(child, size)

			heap.Push(&openQueue, s)
		}
		// FAILURE condition?
	}
}
