package solver

import (
	"math"
	g "n-puzzle/golib"
)

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
