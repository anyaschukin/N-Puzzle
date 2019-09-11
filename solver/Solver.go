package solver

import (
	"bytes"
	"container/heap"
	"fmt"
	g "n-puzzle/golib"
	"os"
	"time"

	// "github.com/AndreasBriese/bbloom"
	// "time"
)

type Problem struct {
	start []int
	goal  []int
	//heuristic      string
	//searchAlgo     string
	solutionPath   map[int][]int // maybe unnecessary?
	sizeComplexity int
	timeComplexity int
	solutionFound  bool
}

func newProblem(Puzzle []int, size int) Problem {
	problem := Problem{}
	problem.start = Puzzle
	problem.goal = MakeGoal(size)
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
	before    *State
	path      *State
}

func newState(Puzzle []int, priority int, depth int, heuristic int, before *State) *State {
	state := &State{}
	state.index = 0
	state.priority = priority
	state.depth = depth         // not sure if we need to store this?
	state.heuristic = heuristic // not sure about this one either?
	state.puzzle = Puzzle
	state.before = before
	state.path = nil
	return state
}

func printPath(state *State, child []int, size int) {
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

func Solver(Puzzle []int, size int) {
	// TESTING RUNTIME
	start := time.Now()

	problem := newProblem(Puzzle, size)
	goal := g.PuzzleToString(problem.goal, ",")

	if IsSolvable(problem.goal, Puzzle, size) == false {
		elapsed := time.Since(start)
		fmt.Println("This puzzle is unsolvable.")
		fmt.Printf("Binomial took %s\n", elapsed)
		os.Exit(1)
	}

	// state := newState(Puzzle, 100000, 0, 0, nil)
	state := newState(Puzzle, 0, 0, 0, nil)

	openSet := make(map[string]int)
	parent := g.PuzzleToString(state.puzzle, ",")
	openSet[parent] = state.priority

	openQueue := CreateQueue(*state)
	// closedSet := bbloom.New(float64(1<<16), float64(0.01))
	closedSet := make(map[string]int)

	unsolved := true
	for unsolved {
		tmp := len(openQueue)
		if tmp > problem.sizeComplexity {
			problem.sizeComplexity = tmp
		}
		if tmp == 0 {
			fmt.Println("This priorityQueue is empty.")
			g.PrintBoard(state.puzzle, size)
			os.Exit(1) //
		}
		state = heap.Pop(&openQueue).(*State)
		parent = g.PuzzleToString(state.puzzle, ",")
		delete(openSet, parent)
		children := CreateNeighbors(state.puzzle, size)

		for _, child := range children {
			tmpChild := g.PuzzleToString(child, ",")

			if bytes.Equal([]byte(goal), []byte(tmpChild)) {
				elapsed := time.Since(start)
				fmt.Println("This puzzle has been solved!\n")
				printPath(state, child, size)
				// Print Space and Time Complexity, & Runtime
				fmt.Printf("Size Complexity: %d\n", problem.sizeComplexity)
				fmt.Printf("Time Complexity: %d\n", problem.timeComplexity)
				fmt.Printf("Binomial took %s", elapsed)
				unsolved = false
				// exit program
				continue
			}

			depth := state.depth - 1
			// depth = -depth
			heuristic := g.Manhattan(child, problem.goal, size)
			s := newState(child, (depth*-1)+heuristic, depth, heuristic, state)

			if _, exists := openSet[tmpChild]; exists {
				if openSet[tmpChild] < s.priority {
					continue
				}
			}

			if _, exists := closedSet[tmpChild]; exists {
				continue
			}
			// if closedSet.Has([]byte(tmpChild)) {	
			// 	continue
			// } 

			openSet[tmpChild] = s.priority
			heap.Push(&openQueue, s)
			problem.timeComplexity++
		}
		closedSet[parent] = state.priority
		// closedSet.AddIfNotHas([]byte(parent))
	}
}
