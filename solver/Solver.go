package solver

import (
	"bytes"
	"container/heap"
	"fmt"
	g "n-puzzle/golib"
	"time"
)

type Problem struct {
	start          []int
	goal           []int
	heuristic      string
	sizeComplexity int
	timeComplexity int
}

func newProblem(Puzzle []int, size int, h string) Problem {
	problem := Problem{}
	problem.start = Puzzle
	problem.goal = MakeGoal(size)
	problem.heuristic = h
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
	state.depth = depth
	state.heuristic = heuristic
	state.puzzle = Puzzle
	state.before = before
	state.path = nil
	return state
}

func Solver(Puzzle []int, size int, h string) {
	start := time.Now()

	problem := newProblem(Puzzle, size, h)
	goal := g.PuzzleToString(problem.goal, ",")

	if IsSolvable(problem.goal, Puzzle, size) == false {
		elapsed := time.Since(start)
		fmt.Println("This puzzle is unsolvable.")
		fmt.Printf("Binomial took %s\n", elapsed)
		return
	}

	state := newState(Puzzle, 0, 0, 0, nil)

	openSet := make(map[string]int)
	parent := g.PuzzleToString(state.puzzle, ",")
	openSet[parent] = state.priority

	openQueue := CreateQueue(*state)
	closedSet := make(map[string]int)

	unsolved := true
	for unsolved {
		tmp := len(openQueue)
		if tmp > problem.sizeComplexity {
			problem.sizeComplexity = tmp
		}
		state = heap.Pop(&openQueue).(*State)
		parent = g.PuzzleToString(state.puzzle, ",")
		delete(openSet, parent)
		children := CreateNeighbors(state.puzzle, size)

		for _, child := range children {
			tmpChild := g.PuzzleToString(child, ",")

			if bytes.Equal([]byte(goal), []byte(tmpChild)) {
				elapsed := time.Since(start)
				PrintSolved(elapsed, problem, state, child, size)
				return
			}

			depth := state.depth - 1
			heuristic := pickHeuristic(child, problem.goal, size, problem.heuristic)
			s := newState(child, (depth*-1)+heuristic, depth, heuristic, state)

			if _, exists := openSet[tmpChild]; exists {
				if openSet[tmpChild] < s.priority {
					continue
				}
			}

			if _, exists := closedSet[tmpChild]; exists {
				continue
			}

			openSet[tmpChild] = s.priority
			heap.Push(&openQueue, s)
			problem.timeComplexity++
		}
		closedSet[parent] = state.priority
	}
}
