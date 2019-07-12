package solver

import (
	"container/heap"
	"fmt"
	p "n-puzzle/parsing"
)

type Item struct {
	index    int
	priority int
	value    []int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value []int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func CreateQueue() {
	Puzzle1 := p.GenerateRandomBoard(3, false, 0)
	fmt.Printf("Puzzle 1 = %v\n", Puzzle1)
	Puzzle2 := p.GenerateRandomBoard(3, false, 0)
	fmt.Printf("Puzzle 2 = %v\n", Puzzle2)
	Puzzle3 := p.GenerateRandomBoard(3, false, 0)
	fmt.Printf("Puzzle 3 = %v\n", Puzzle3)

	items := map[int][]int{
		1: Puzzle1, 2: Puzzle2, 3: Puzzle3,
	}

	pq := make(PriorityQueue, len(items))
	i := 0
	for priority, value := range items {
		pq[i] = &Item{
			index:    i,
			priority: priority,
			value:    value,
		}
		i++
	}
	heap.Init(&pq)

	Puzzle4 := p.GenerateRandomBoard(3, false, 0)
	fmt.Printf("Puzzle 4 = %v\n", Puzzle4)
	item := &Item{
		priority: 1,
		value:    Puzzle4,
	}
	heap.Push(&pq, item)
	pq.update(item, item.value, 5)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%v \n", item.priority, item.value)
	}
}
