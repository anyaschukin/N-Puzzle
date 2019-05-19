package golib

func MakeRangeNum(min, max int) []int {
	Set := make([]int, max-min+1)
	for i := range Set {
		Set[i] = min + i
	}
	return Set
}
