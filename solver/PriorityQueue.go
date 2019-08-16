package solver

import (
	"container/heap"
)

//type State struct {
//	index     int
//	priority  int
//	cost      int
//	heuristic int
//	puzzle    []int
//}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	state := x.(*State)
	state.index = n
	*pq = append(*pq, state)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	state := old[n-1]
	state.index = -1 // for safety
	*pq = old[0 : n-1]
	return state
}

func (pq *PriorityQueue) Update(state *State, puzzle []int, priority int) {
	state.puzzle = puzzle
	state.priority = priority
	heap.Fix(pq, state.index)
}

func CreateQueue(state State) PriorityQueue {
	Pq := make(PriorityQueue, 1)
	Pq[0] = &state
	heap.Init(&Pq)
	return Pq
}

//	for priority, puzzle := range States {
//		pq[i] = &State{
//			index:    i,
//			priority: priority,
//			puzzle:    puzzle,
//		}
//		i++
//	}

//		Puzzle4 := p.GenerateRandomBoard(3)
//		fmt.Printf("Puzzle 4 = %v\n", Puzzle4)
// state := &State{
// 	priority: 1,
// 	puzzle:    Puzzle4,
// }
// heap.Push(&pq, state)
// pq.update(state, state.puzzle, 5)
//
//		for pq.Len() > 0 {
//			State := heap.Pop(&pq).(*State)
//			fmt.Printf("%.2d:%v \n", State.priority, State.puzzle)
//		}
