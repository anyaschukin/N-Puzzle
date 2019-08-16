package solver

import (
	"container/heap"
	"fmt"
	g "n-puzzle/golib"
	"os"
	"reflect"

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

// if I comment this
func Solver(Puzzle []int, size int, iterations int) {
	problem := newProblem(Puzzle, size)

	if IsSolvable(problem.goal, Puzzle, size) == false {
		fmt.Println("This puzzle in unsolvable.")
		os.Exit(1)
	}
	closedSet := bbloom.New(float64(1<<16), float64(0.01))
	state := newState(Puzzle, size, 1, 0, 0)
	openQueue := CreateQueue(*state)

	for counter := 0; counter < 300000; counter++ {
		state = heap.Pop(&openQueue).(*State)

		if reflect.DeepEqual(problem.goal, state.puzzle) {
			fmt.Println("This puzzle has been solved!\n")
			os.Exit(1)
		}

		fmt.Println("\nNODE\n")

		fmt.Printf("goal: %v \n", problem.goal)
		fmt.Printf("%d, %d, %d: %v \n", state.priority, state.depth, state.heuristic, state.puzzle)
		//time.Sleep(1 * time.Second)
		closedSet.Add([]byte(g.PuzzleToString(state.puzzle, ",")))

		children := CreateNeighbors(state.puzzle, size)
		for _, child := range children {
			if closedSet.Has([]byte(g.PuzzleToString(child, ","))) {
				problem.sizeComplexity++
				continue
			} else {
				problem.timeComplexity++
				heuristic := g.Manhattan(child, problem.goal, size)
				priority := (state.depth + 1) + heuristic
				s := newState(child, size, priority, state.depth+1, heuristic)
				fmt.Printf("%d, %d, %d: %v \n", s.priority, s.depth, s.heuristic, s.puzzle)

				//closedSet.Add([]byte(g.PuzzleToString(child, ",")))
				heap.Push(&openQueue, s)
				fmt.Printf("priority = %d\n", priority)
			}
			//			tentative_gScore := g.Manhattan(state.puzzle, child, size)
			//			gScore := g.Manhattan(child, problem.goal, size)

			// if s in target... rebuild path to start and finish!
			//heap.Push((&openQueue).(*State))
			//heap.Push((&openQueue).(*State))

		}
	}

}

// Puzzle4 := p.GenerateRandomBoard(3)
// state = newState(Puzzle4, size, 0)
// heap.Push(&openQueue, state)
// openQueue.Update(state, state.puzzle, 4)

// Puzzle3 := p.GenerateRandomBoard(3)
// state = newState(Puzzle3, size, 0)
// heap.Push(&openQueue, state)
// openQueue.Update(state, state.puzzle, 3)

// Puzzle2 := p.GenerateRandomBoard(3)
// state = newState(Puzzle2, size, 0)
// heap.Push(&openQueue, state)
// openQueue.Update(state, state.puzzle, 2)

// Puzzle5 := p.GenerateRandomBoard(3)
// state = newState(Puzzle5, size, 0)
// heap.Push(&openQueue, state)
// openQueue.Update(state, state.puzzle, 5)

// for openQueue.Len() > 0 {
// 	state := heap.Pop(&openQueue).(*State)
// 	fmt.Printf("%.2d:%v \n", state.priority, state.puzzle)
// }

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
