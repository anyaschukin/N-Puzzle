package golib

import "math/rand"

// returns a random boolean

func RandomBool() bool {
	return rand.Float32() < 0.5
}
