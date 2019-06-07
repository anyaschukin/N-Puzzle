package golib

import (
	"fmt"
)

func PrintBoard(slice []int, size int) {
	i := 0
	fmt.Print("\n\n")
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if slice[i] < 10 {
				fmt.Printf("%v   ", slice[i])
			} else {
				fmt.Printf("%v  ", slice[i])
			}
			i++
		}
		fmt.Print("\n")
	}
}
