package golib

func SpiralMatrix(board []int, size int) []int {
	left, top, right, bottom := 0, 0, size-1, size-1
	puzzle := make([]int, size*size)
	cur := 0
	for left < right {
		// work right, along top
		for i := left; i <= right; i++ {
			puzzle[top*size+i] = board[cur]
			cur++
		}
		top++
		// work down right side
		for j := top; j <= bottom; j++ {
			puzzle[j*size+right] = board[cur]
			cur++
		}
		right--
		if top == bottom {
			break
		}
		// work left, along bottom
		for i := right; i >= left; i-- {
			puzzle[bottom*size+i] = board[cur]
			cur++
		}
		bottom--
		// work up left size
		for j := bottom; j >= top; j-- {
			puzzle[j*size+left] = board[cur]
			cur++
		}
		left++
	}
	// center (last) element
	puzzle[top*size+left] = board[cur]
	return puzzle
}
