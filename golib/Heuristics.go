package golib

// Hamming returns the sum of misplaced tiles.
func Hamming(board []int, target []int, s int) int {
	length := s * s
	hamming := 0
	for i := 0; i < length; i++ {
		if board[i] != target[i] && board[i] != 0 {
			hamming++
		}
	}
	return hamming
}

func indexToCoordinates(index int, size int) (int, int) {
	x := index % size
	y := index / size
	return x, y
}

func absInt(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

// ManhattanDistance returns the manhattan distance between two misplaced tiles
func ManhattanDistance(val int, indexBoard int, size int, target []int) int {
	indexTarget := FindIndexSlice(target, val)
	xT, yT := indexToCoordinates(indexTarget, size)
	xC, yC := indexToCoordinates(indexBoard, size)
	return (absInt(xT-xC) + absInt(yT-yC))
}

// Manhattan returns the sum of manhattan distances .
func Manhattan(board []int, target []int, s int) int {
	length := s * s
	manhattan := 0
	for i := 0; i < length; i++ {
		if board[i] != target[i] && board[i] != 0 {
			manhattan += ManhattanDistance(board[i], i, s, target)
		}
	}
	return manhattan
}
