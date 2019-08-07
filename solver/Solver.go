package solver

import "fmt"

type Solve struct {
	start          []int
	goal           []int
	solvable       bool
	heuristic      string
	searchAlgo     string
	weight         []int
	solutionPath   []int
	sizeComplexity int
	timeComplexity int
	solutionFound  bool
}

func newSolve(Puzzle []int, size int) Solve {
	solve := Solve{}
	solve.start = Puzzle
	solve.goal = MakeGoal(size)
	solve.solvable = IsSolvable(solve.goal, Puzzle, size)
	solve.heuristic = "MANHATTAN"
	solve.searchAlgo = "A_STAR"
	solve.sizeComplexity = 0
	solve.timeComplexity = 1
	solve.solutionFound = false
	return solve
}

// if I comment this
func Solver(Puzzle []int, size int, iterations int) {
	solve := newSolve(Puzzle, size)
	counter := 0
	closedList := 0
	closedSet := 

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
}

//for i := 1; i <= iterations; i++ {
//	swapEmpty(p, size)
//}
