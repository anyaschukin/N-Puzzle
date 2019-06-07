package golib

import (
	"math/rand"
	"time"
)

func RandomChoice(Poss []int) int {
	rand.Seed(time.Now().Unix())
	swi := rand.Intn(len(Poss))
	return swi
}
