package solver

import (
	//"fmt"
	"container/heap"
)

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
