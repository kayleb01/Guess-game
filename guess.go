package main

import (
	"math/rand"
	"time"
)

func guesses() int {

	var min int = 1
	var max int = 200
	rand.Seed(time.Now().UnixNano())
	RamdomInt := rand.Intn(max-min) + min

	return RamdomInt
}
