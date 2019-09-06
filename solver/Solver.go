package solver

import (
	"bytes"
	"container/heap"
	"fmt"
	"log"
	g "n-puzzle/golib"
	"os"
<<<<<<< HEAD
	"reflect"
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> testing speed reflect.DeepEqual vs bytes.Equal
	"time"

=======
	// "time"
>>>>>>> priorityqueue empty before solution
=======

>>>>>>> new openSet same problems
	"github.com/AndreasBriese/bbloom"
)

type Problem struct {
	start []int
	goal  []int
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
	// solvable bool
>>>>>>> new openSet same problems
=======
>>>>>>> solves half of time for size 7 + 8
=======
>>>>>>> little commit before checkout
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
<<<<<<< HEAD
<<<<<<< HEAD
=======
	// problem.solvable = IsSolvable(problem.goal, Puzzle, size)
>>>>>>> new openSet same problems
=======
>>>>>>> little commit before checkout
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
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
func newState(Puzzle []int, priority int, depth int, heuristic int, before *State) *State {
	state := &State{}
	state.index = 0
	state.priority = priority
	state.depth = depth         // not sure if we need to store this?
	state.heuristic = heuristic // not sure about this one either?
<<<<<<< HEAD
=======
func newState(Puzzle []int, priority int, depth int, heuristic int) *State {
=======
func newState(Puzzle []int, priority int, depth int, heuristic int, before *State) *State {
>>>>>>> builds path from finish -> start
=======
func newState(Puzzle []int, priority int, depth int, heuristic int, before *State, after *State) *State {
>>>>>>> new binary for Drew tests
=======
func newState(Puzzle []int, priority int, depth int, heuristic int, before *State) *State {
>>>>>>> small edits to Sovler
	state := &State{}
	state.index = 0
	state.priority = priority
<<<<<<< HEAD
	state.depth = depth         // not sure if I need to store this either?
	state.heuristic = heuristic // I don't think I need to keep this?
>>>>>>> new openSet same problems
=======
	state.depth = depth         // not sure if we need to store this?
	state.heuristic = heuristic // not sure about this one either?
>>>>>>> solves half of time for size 7 + 8
=======
>>>>>>> little commit before checkout
	state.puzzle = Puzzle
	state.before = before
	return state
}

func Solver(Puzzle []int, size int) {
	// TESTING RUNTIME
	start := time.Now()
<<<<<<< HEAD
=======
func Solver(Puzzle []int, size int, iterations int) {
	start := time.Now()

	problem := newProblem(Puzzle, size)
	unsolved := true
>>>>>>> new openSet same problems

<<<<<<< HEAD
	problem := newProblem(Puzzle, size)
	goal := g.PuzzleToString(problem.goal, ",")
=======
	goal := g.PuzzleToString(problem.goal, ",")
	// g.PrintBoard(Puzzle, size)
>>>>>>> testing speed reflect.DeepEqual vs bytes.Equal
=======

	problem := newProblem(Puzzle, size)
	goal := g.PuzzleToString(problem.goal, ",")
>>>>>>> little commit before checkout

	if IsSolvable(problem.goal, Puzzle, size) == false {
		fmt.Println("This puzzle is unsolvable.")
		os.Exit(1)
	}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	state := newState(Puzzle, 100000, 0, 0, nil)
=======
	state := newState(Puzzle, 100000, 0, 0)
>>>>>>> new openSet same problems
=======
	state := newState(Puzzle, 100000, 0, 0, nil)
>>>>>>> builds path from finish -> start
=======
	state := newState(Puzzle, 100000, 0, 0, nil, nil)
>>>>>>> new binary for Drew tests
=======
	state := newState(Puzzle, 100000, 0, 0, nil)
>>>>>>> small edits to Sovler

	openSet := make(map[string]int)
	parent := g.PuzzleToString(state.puzzle, ",")
	openSet[parent] = state.priority

<<<<<<< HEAD
<<<<<<< HEAD
=======
	// closedSet := make(map[string]int)
	closedSet := bbloom.New(float64(1<<16), float64(0.01))
>>>>>>> new openSet same problems
	openQueue := CreateQueue(*state)
	closedSet := bbloom.New(float64(1<<16), float64(0.01))

<<<<<<< HEAD
<<<<<<< HEAD
	unsolved := true
	for unsolved {

=======
	for counter := 0; counter < 6000000; counter++ {
		
>>>>>>> priorityqueue empty before solution
=======
=======
	openQueue := CreateQueue(*state)
	closedSet := bbloom.New(float64(1<<16), float64(0.01))

	unsolved := true
>>>>>>> little commit before checkout
	for unsolved {

>>>>>>> new openSet same problems
		if len(openQueue) == 0 {
			fmt.Println("This priorityQueue is empty.")
			g.PrintBoard(state.puzzle, size)
			os.Exit(1)
		}

		state = heap.Pop(&openQueue).(*State)
		parent = g.PuzzleToString(state.puzzle, ",")
		delete(openSet, parent)

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
		closedSet.AddIfNotHas([]byte(parent))
=======
		// fmt.Println(" ----------- ")
		// fmt.Println("\n NEW STATE")
		// fmt.Printf("\n priority = %d, heuristic = %d, depth = %d", state.priority, state.heuristic, state.depth)
		// g.PrintBoard(state.puzzle, size)
		// fmt.Println(" ----------- ")
>>>>>>> priorityqueue empty before solution
=======
		closedSet.AddIfNotHas([]byte(parent))
		// if _, exists := closedSet[parent]; !exists {
		// 	closedSet[parent] = state.priority
		// }
>>>>>>> new openSet same problems
=======
		closedSet.AddIfNotHas([]byte(parent))
>>>>>>> little commit before checkout

		if bytes.Equal([]byte(parent), []byte(goal)) {
			fmt.Println("This puzzle has been solved!\n")
			// REBUILD PATH TO START
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
			for p := state; p != nil; p = state.before {
				g.PrintBoard(p.puzzle, size)
				if reflect.DeepEqual(problem.goal, p.puzzle) {
					break
				}
			}
=======
>>>>>>> little commit before checkout
=======
			for p := state; p.before != nil; p = p.before {
				g.PrintBoard(p.puzzle, size)
				time.Sleep(1 * time.Second)
			}
=======
			// for p := state; p.before != nil; p = p.before {

			// 	g.PrintBoard(p.puzzle, size)
			// 	time.Sleep(1 * time.Second)
			// }
>>>>>>> new binary for Drew tests
			g.PrintBoard(problem.start, size)
>>>>>>> builds path from finish -> start

			// TESTING RUNTIME
			elapsed := time.Since(start)
			fmt.Printf("Binomial took %s", elapsed)
			unsolved = false
			// os.Exit(1)
<<<<<<< HEAD
		}

<<<<<<< HEAD
		// if reflect.DeepEqual(problem.goal, state.puzzle) {
		// fmt.Println("This puzzle has been solved!\n")
		// g.PrintBoard(state.puzzle, size)
		// REBUILD PATH TO START
		// unsolved = false
		// elapsed := time.Since(start)
		// log.Printf("Binomial took %s", elapsed)
		// os.Exit(1)
		// }

<<<<<<< HEAD
=======
>>>>>>> new binary for Drew tests
		children := CreateNeighbors(state.puzzle, size)
<<<<<<< HEAD
=======
		// fmt.Println("--- children ---")
		// fmt.Print(children)
>>>>>>> priorityqueue empty before solution
=======
			unsolved = false
=======
			elapsed := time.Since(start)
			log.Printf("Binomial took %s", elapsed)
>>>>>>> testing speed reflect.DeepEqual vs bytes.Equal
			os.Exit(1)
=======
>>>>>>> little commit before checkout
		}

		// if reflect.DeepEqual(problem.goal, state.puzzle) {
		// fmt.Println("This puzzle has been solved!\n")
		// g.PrintBoard(state.puzzle, size)
		// REBUILD PATH TO START
		// unsolved = false
		// elapsed := time.Since(start)
		// log.Printf("Binomial took %s", elapsed)
		// os.Exit(1)
		// }
<<<<<<< HEAD

		// fmt.Printf("\n-- parent --")
		// g.PrintBoard(state.puzzle, size)
		// fmt.Printf("\n priority = %d, heuristic = %d, depth = %d\n", state.priority, state.heuristic, state.depth)

		// time.Sleep(1000 * time.Millisecond)
=======
>>>>>>> little commit before checkout

=======
>>>>>>> little test
		children := CreateNeighbors(state.puzzle, size)
<<<<<<< HEAD
		// fmt.Printf("\n-- child --")
>>>>>>> new openSet same problems
=======
>>>>>>> little commit before checkout

		for _, child := range children {
			tmpChild := g.PuzzleToString(child, ",")

			if bytes.Equal([]byte(parent), []byte(tmpChild)) {
				fmt.Println("This puzzle has been solved!\n")
				g.PrintBoard(state.puzzle, size)
				// REBUILD PATH TO START
<<<<<<< HEAD
				// THIS DOESNT WORK YET
=======
>>>>>>> builds path from finish -> start
				// for p := state; p != nil; p = state.before {
				// 	g.PrintBoard(state.puzzle, size)
				// }

				// TESTING RUNTIME
				elapsed := time.Since(start)
				fmt.Printf("Binomial took %s", elapsed)
				unsolved = false
			}

<<<<<<< HEAD
			// if reflect.DeepEqual(problem.goal, child) {
			// fmt.Println("This puzzle has been solved!\n")
			// g.PrintBoard(child, size)
<<<<<<< HEAD
<<<<<<< HEAD
			// REBUILD PATH TO START
			// elapsed := time.Since(start)
			// log.Printf("Binomial took %s", elapsed)
			// os.Exit(1)
			// }

<<<<<<< HEAD
=======
>>>>>>> new binary for Drew tests
			if closedSet.Has([]byte(tmpChild)) {
=======
			if reflect.DeepEqual(problem.goal, child) {
=======
			tmpChild := g.PuzzleToString(child, ",")

			if bytes.Equal([]byte(parent), []byte(tmpChild)) {
>>>>>>> testing speed reflect.DeepEqual vs bytes.Equal
				fmt.Println("This puzzle has been solved!\n")
				g.PrintBoard(state.puzzle, size)
				// REBUILD PATH TO START
				elapsed := time.Since(start)
				log.Printf("Binomial took %s", elapsed)
				os.Exit(1)
			}

			// if reflect.DeepEqual(problem.goal, child) {
			// fmt.Println("This puzzle has been solved!\n")
			// g.PrintBoard(child, size)
=======
>>>>>>> little commit before checkout
			// REBUILD PATH TO START
			// elapsed := time.Since(start)
			// log.Printf("Binomial took %s", elapsed)
			// os.Exit(1)
			// }

			if closedSet.Has([]byte(g.PuzzleToString(child, ","))) {
>>>>>>> new openSet same problems
				problem.sizeComplexity++
				continue
			}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> little commit before checkout
			depth := -(state.depth + 1)
			// depth = -depth
			heuristic := g.Manhattan(child, problem.goal, size)
<<<<<<< HEAD
			s := newState(child, depth+heuristic, depth, heuristic, state)

<<<<<<< HEAD
			if _, exists := openSet[tmpChild]; exists {
				if openSet[tmpChild] < s.priority {
					continue
				}
			}
=======
			priority := state.depth + 1 + heuristic
=======
			depth := state.depth + 1
			depth = -depth
=======
			depth := -(state.depth + 1)
			// depth = -depth
>>>>>>> solves half of time for size 7 + 8
			heuristic := g.Manhattan(child, problem.goal, size)
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>> new openSet same problems
			// priority = -priority
			// priority :=  heuristic
			s := newState(child, depth+heuristic, depth, heuristic)
=======
			s := newState(child, depth+heuristic, depth, heuristic, state)
>>>>>>> builds path from finish -> start

			// if _, exists := closedSet[tmpChild]; exists {
			// 	continue
			// }
			// g.PrintBoard(child, size)
<<<<<<< HEAD
>>>>>>> priorityqueue empty before solution
=======
			// fmt.Printf("\n priority = %d, heuristic = %d, depth = %d\n", s.priority, s.heuristic, s.depth)
=======
			s := newState(child, depth+heuristic, depth, heuristic, state, nil)
>>>>>>> new binary for Drew tests
=======
			s := newState(child, depth+heuristic, depth, heuristic, state)
>>>>>>> small edits to Sovler

=======
>>>>>>> little commit before checkout
			if _, exists := openSet[tmpChild]; exists {
				if openSet[tmpChild] < s.priority {
					continue
				}
			}
>>>>>>> new openSet same problems

			openSet[tmpChild] = s.priority
			heap.Push(&openQueue, s)

			// problem.timeComplexity++

<<<<<<< HEAD
<<<<<<< HEAD
=======
			// s := newState(child, size, priority, state.depth+1, heuristic)

>>>>>>> new openSet same problems
		}
	}
}

// if bytes.Equal([]byte(g.PuzzleToString(problem.goal, ",")), []byte(g.PuzzleToString(state.puzzle, ","))) {
// 	fmt.Println("This puzzle has been solved!\n")
// 	g.PrintBoard(state.puzzle, size)
// 	// REBUILD PATH TO START
// 	os.Exit(1)
// }

// fmt.Println(" ----------- ")
// fmt.Println("\n NEW STATE")
// fmt.Printf("\n priority = %d, heuristic = %d, depth = %d", state.priority, state.heuristic, state.depth)
// g.PrintBoard(state.puzzle, size)
// fmt.Println(" ----------- ")

// if s.priority > state.priority {
// 	continue
// }
// fmt.Printf("\n priority = %d, heuristic = %d, depth = %d", priority, heuristic, state.depth + 1)
// g.PrintBoard(child, size)
=======
		}
	}
}
<<<<<<< HEAD
>>>>>>> little commit before checkout
=======

// if reflect.DeepEqual(problem.goal, state.puzzle) {
// fmt.Println("This puzzle has been solved!\n")
// g.PrintBoard(state.puzzle, size)
// REBUILD PATH TO START
// unsolved = false
// elapsed := time.Since(start)
// log.Printf("Binomial took %s", elapsed)
// os.Exit(1)
// }
>>>>>>> new binary for Drew tests
