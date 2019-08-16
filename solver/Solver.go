package solver

import (
	"container/heap"
	"fmt"
	p "n-puzzle/parsing"
	"os"
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

//func (problem *Problem) solution(puzzle []int) {
//	if len(a) != len(b) {
//		Problem.solutionFound = false
//		return false
//	}
//	for i := range a {
//		if a[i] != b[i] {
//			Problem.solutionFound = false
//			return false
//		}
//	}
//	Problem.solutionFound = true
//	return true
//}

type State struct {
	index     int
	priority  int
	depth     int
	heuristic int
	puzzle    []int
}

func newState(Puzzle []int, size int, depth int) *State {
	state := &State{}
	state.index = 0
	state.priority = 1
	state.depth = depth
	state.heuristic = 0
	state.puzzle = Puzzle
	return state
}

// if I comment this
func Solver(Puzzle []int, size int, iterations int) {
	problem := newProblem(Puzzle, size)
	if IsSolvable(problem.goal, Puzzle, size) == false {
		fmt.Println("This puzzle in unsolvable.")
		os.Exit(1)
	}

	//closedSet := bbloom.New(float64(1<<16), float64(0.01))
	state := newState(Puzzle, size, 0)
	openQueue := CreateQueue(*state)
	//state = heap.Pop(&openQueue).(*State)

	Puzzle4 := p.GenerateRandomBoard(3)
	state = newState(Puzzle4, size, 0)
	heap.Push(&openQueue, state)
	openQueue.Update(state, state.puzzle, 4)

	Puzzle3 := p.GenerateRandomBoard(3)
	state = newState(Puzzle3, size, 0)
	heap.Push(&openQueue, state)
	openQueue.Update(state, state.puzzle, 3)

	for openQueue.Len() > 0 {
		state := heap.Pop(&openQueue).(*State)
		fmt.Printf("%.2d:%v \n", state.priority, state.puzzle)
	}

	//for counter := 0; counter < 300000; counter++ {
	//	state = heap.Pop(&openQueue).(*State)
	//	closedSet.Add([]byte(g.PuzzleToString(state.puzzle, ",")))
	//	if problem.solution(Puzzle) {
	//
	//	}
	//	for child := range CreateNeighbors(state.Puzzle, size) {
	//		s = newState(child, size, state.depth+1)
	//		// if s in target... rebuild path to start and finish!
	//		if closedset.Has([]byte(g.PuzzleToString(s.puzzle, ","))) {
	//			problem.sizeComplexity++
	//		} else {
	//			problem.timeComplexity++
	//			heap.Push((&openQueue).(*State))
	//		}
	//
	//	}
	//}
}

//Solution := MakeGoal(size)
//g.PrintBoard(Solution, size)
//
//solveable := IsSolvable(Solution, Puzzle, size)
//fmt.Printf("\nsolve it? %v\n", solveable)

//CreateQueue()

//neighbors := CreateNeighbors(Puzzle, size)
//fmt.Println(neighbors)

// apply algo:
// 	heuristic flags()
// 	switch case algo1 algo2 algo3
// 		build priority queue
//		generate neighbors, explore, build history
// 		solved?

//goal := g.CheckSliceEquality(Puzzle, Solution)
//fmt.Printf("goal? %v\n", goal)

//for i := 1; i <= iterations; i++ {
//	swapEmpty(p, size)
//}
