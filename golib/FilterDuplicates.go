package golib

func FilterDuplicates(input []int) []int {
	Unique := make([]int, 0, len(input))
	check := make(map[int]bool)

	for _, val := range input {
		if _, ok := check[val]; !ok {
			check[val] = true
			Unique = append(Unique, val)
		}
	}
	return Unique
}
