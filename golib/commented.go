package golib

//func swapEmpty(puzzle []int, size int) {
//	//fmt.Print(puzzle)e
//	idx := g.FindIndexSlice(puzzle, 0)
//	//fmt.Print(idx)
//	//fmt.Println("\n")
//	var Poss []int
//	if idx%size > 0 {
//		Poss = append(Poss, idx-1)
//	}
//	if idx%size < size-1 {
//		Poss = append(Poss, idx+1)
//	}
//	if idx/size > 0 {
//		Poss = append(Poss, idx-size)
//	}
//	if idx/size < size-1 {
//		Poss = append(Poss, idx+size)
//	}
//	//fmt.Print(Poss)
//	swi := g.RandomChoice(Poss)
//	//fmt.Print(swi)
//	puzzle[idx] = puzzle[swi]
//	puzzle[swi] = 0
//}

// need to change MakeGoal to spiral/snail format
//func makeGoal(size int) []int {
//	totalSize := size * size
//	puzzle := make([]int, totalSize)
//	cur := 1
//	for i := range puzzle {
//		if i == (size*size)-1 {
//			puzzle[i] = 0
//			break
//		}
//		puzzle[i] = cur
//		cur++
//	}
//	fmt.Printf("goal = ")
//	g.PrintBoard(puzzle, size)
//	return puzzle
//}
