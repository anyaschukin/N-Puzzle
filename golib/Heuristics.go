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
func manhattanDistance(val int, indexBoard int, size int, target []int) int {
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
			manhattan += manhattanDistance(board[i], i, s, target)
		}
	}
	return manhattan
}

// outRow returns true if tile is out of Row
func outRow(board []int, target []int, s int, length int, i int) bool {
	for col := 0; col < s; col++ {
		if board[i] == target[i-(i%s)+col] {
			return false
		}
	}
	return true
}

// outCol returns true if tile is out of Column
func outCol(board []int, target []int, s int, length int, i int) bool {
	for row := 0; row < s; row++ {
		if board[i] == target[(i%s)+(row*s)] {
			return false
		}
	}
	return true
}

// OutRowCol returns Number of tiles out of row plus number of tiles out of column.
func OutRowCol(board []int, target []int, s int) int {
	length := s * s
	outRowCol := 0
	for i := 0; i < length; i++ {
		if outRow(board, target, s, length, i) && board[i] != 0 {
			outRowCol++
		}
	}
	for i := 0; i < length; i++ {
		if outCol(board, target, s, length, i) && board[i] != 0 {
			outRowCol++
		}
	}
	return outRowCol
}
