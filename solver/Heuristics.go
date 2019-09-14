package solver

import (
	"math"
	g "n-puzzle/golib"
)

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
	indexTarget := g.FindIndexSlice(target, val)
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

// foundCurrent returns true if the position matches the current tile
func foundCurrent(currentRow int, currentCol int, row int, col int) bool {

	if currentRow == row && currentCol == col {
		return true
	}
	return false
}

// findNext returns the string index of next tile in snail order
func findNext(s int, current int) int {
	left, top, right, bottom := 0, 0, s-1, s-1
	row, col := 0, 0
	next := 0
	currentRow := current / s
	currentCol := current % s
	found := false
	first := true
	for left < right {
		// work right, along top
		for i := left; i <= right; i++ {
			if first == true && i == right {
				first = false
				continue
			}
			next++
			found = foundCurrent(currentRow, currentCol, row, col)
			if found {
				return next
			}
			col++
		}
		top++
		// work down right side
		for j := top; j <= bottom; j++ {
			next += s
			found = foundCurrent(currentRow, currentCol, row, col)
			if found {
				return next
			}
			row++
		}
		right--
		if top == bottom {
			continue
		}
		// work left, along bottom
		for i := right; i >= left; i-- {
			next--
			found = foundCurrent(currentRow, currentCol, row, col)
			if found {
				return next
			}
			col--
		}
		bottom--
		// work up left size
		for j := bottom; j >= top; j-- {
			next -= s
			found = foundCurrent(currentRow, currentCol, row, col)
			if found {
				return next
			}
			row--
		}
		left++
	}
	for i := left; i <= right; i++ {
		if first == true && i == right {
			first = false
			continue
		}
		next++
		found = foundCurrent(currentRow, currentCol, row, col)
		if found {
			return next
		}
		col++
	}
	// fmt.Printf("Error in func foundCurrent for Nilsson's Sequence Score. Should we get here?\n")
	return next
}

// Nilsson's Sequence Score: Manhattan + 3 S(n)
func Nilsson(board []int, target []int, s int) int {
	length := s * s
	manhattan := Manhattan(board, target, s)
	sequenceScore := 0
	current := 0
	next := 1
	for i := 1; i < length; i++ {
		next := findNext(s, current)
		if target[next] != board[current]+1 {
			sequenceScore += 2
		}
		current = next
	}
	if board[next] != 0 {
		sequenceScore++
	}
	nilsson := manhattan + 3*sequenceScore
	return nilsson
}

// Euclidean returns the Euclidean distance between misplaced tiles.
func Euclidean(board []int, target []int, s int) int {
	length := s * s
	euclidean := 0.0
	for i := 0; i < length; i++ {
		goal := g.FindIndexSlice(target, board[i])
		euclidean += (math.Pow(float64(goal/s), 2) + math.Pow(float64(goal%s), 2))
	}
	euclidean = math.Round(math.Sqrt(euclidean))
	return int(euclidean)
}

func pickHeuristic(board []int, target []int, size int, heuristic string) int {
	value := 0
	switch heuristic {
	case "hamming":
		value = Hamming(board, target, size)
	case "euclidean":
		value = Euclidean(board, target, size)
	case "nilsson":
		value = Nilsson(board, target, size)
	case "outRowCol":
		value = OutRowCol(board, target, size)
	default:
		// manhattan is default heuristic
		value = Manhattan(board, target, size)
	}
	return value
}
