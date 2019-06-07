package golib

func FindIndexSlice(slice []int, value int) int {
	for p, v := range slice {
		if value == v {
			return p
		}
	}
	return -1
}
