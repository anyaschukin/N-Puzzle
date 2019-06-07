package solver

func MakeGoal(size int) []int {
	left, top, right, bottom := 0, 0, size-1, size-1
	puzzle := make([]int, size*size)
	cur := 1
	for left < right {
		// work right, along top
		for i := left; i <= right; i++ {
			puzzle[top*size+i] = cur
			cur++
		}
		top++
		// work down right side
		for j := top; j <= bottom; j++ {
			puzzle[j*size+right] = cur
			cur++
		}
		right--
		if top == bottom {
			break
		}
		// work left, along bottom
		for i := right; i >= left; i-- {
			puzzle[bottom*size+i] = cur
			cur++
		}
		bottom--
		// work up left size
		for j := bottom; j >= top; j-- {
			puzzle[j*size+left] = cur
			cur++
		}
		left++
	}
	// center (last) element
	puzzle[top*size+left] = 0
	return puzzle
}

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
