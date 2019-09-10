package solver

import (
	"bytes"
	"container/heap"
	"fmt"
	g "n-puzzle/golib"
	"os"
	"time"

	"github.com/AndreasBriese/bbloom"
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

func printPath(state *State, size int) {
	var tmp *State
	tmp = nil
	for p := state; p != nil; p = p.before {
		p.path = tmp
		tmp = p
	}
	for p := tmp; p != nil; p = p.path {
		g.PrintBoard(p.puzzle, size)
	}
}

func Solver(Puzzle []int, size int) {
	// TESTING RUNTIME
	start := time.Now()

	problem := newProblem(Puzzle, size)
	goal := g.PuzzleToString(problem.goal, ",")

	if IsSolvable(problem.goal, Puzzle, size) == false {
		fmt.Println("This puzzle is unsolvable.")
		os.Exit(1)
	}

	state := newState(Puzzle, 100000, 0, 0, nil)

	openSet := make(map[string]int)
	parent := g.PuzzleToString(state.puzzle, ",")
	openSet[parent] = state.priority

	openQueue := CreateQueue(*state)
	closedSet := bbloom.New(float64(1<<16), float64(0.01))

	unsolved := true
	for unsolved {

		if len(openQueue) == 0 {
			fmt.Println("This priorityQueue is empty.")
			g.PrintBoard(state.puzzle, size)
			os.Exit(1)
		}

		state = heap.Pop(&openQueue).(*State)
		parent = g.PuzzleToString(state.puzzle, ",")
		delete(openSet, parent)

		closedSet.AddIfNotHas([]byte(parent))

		if bytes.Equal([]byte(parent), []byte(goal)) {
			fmt.Println("This puzzle has been solved!\n")
			g.PrintBoard(state.puzzle, size)
			// REBUILD PATH TO START
			// for p := state; p != nil; p = p.before {
			// 	g.PrintBoard(p.puzzle, size)
			// 	if reflect.DeepEqual(problem.goal, p.puzzle) {
			// 		break
			// 	}
			// }

			// TESTING RUNTIME
			elapsed := time.Since(start)
			fmt.Printf("Binomial took %s", elapsed)
			unsolved = false
			// os.Exit(1)
		}

		children := CreateNeighbors(state.puzzle, size)

		for _, child := range children {
			tmpChild := g.PuzzleToString(child, ",")

			if bytes.Equal([]byte(goal), []byte(tmpChild)) {
				fmt.Println("This puzzle has been solved!\n")
				printPath(state, size)
				g.PrintBoard(child, size)

				// TESTING RUNTIME
				elapsed := time.Since(start)
				fmt.Printf("Binomial took %s", elapsed)
				unsolved = false
			}

			if closedSet.Has([]byte(tmpChild)) {
				problem.sizeComplexity++
				continue
			}

			depth := -(state.depth + 1)
			// depth = -depth
			heuristic := g.Manhattan(child, problem.goal, size)
			s := newState(child, depth+heuristic, depth, heuristic, state)

			if _, exists := openSet[tmpChild]; exists {
				if openSet[tmpChild] < s.priority {
					continue
				}
			}

			openSet[tmpChild] = s.priority
			heap.Push(&openQueue, s)

			// problem.timeComplexity++

		}
	}
}

// if reflect.DeepEqual(problem.goal, child) {
// fmt.Println("This puzzle has been solved!\n")
// g.PrintBoard(child, size)
// REBUILD PATH TO START
// elapsed := time.Since(start)
// log.Printf("Binomial took %s", elapsed)
// os.Exit(1)
// }
