package main

import (
	"fmt"
	"math/rand/v2"
)

// use math/rand/v2 for pseudorandom number generation

func main() {
	// rand.IntN for a randon int between 0 and x
	p := fmt.Println
	p(rand.IntN(100))

	// rand.Float64 for a number between 0.0 and 1.0
	p(rand.Float64())

	// example of a float64 between 5.0 and 10.0
	p((rand.Float64() * 5.0) + 5.0)

	// ----- USE SEED IN GENERATORS --------

	seed := rand.NewPCG(42, 1024) // create a rand.Source struct (permuted congruential generator)
	random := rand.New(seed)

	p(random.IntN(100))

	// if you create another seed using the same 42 and 1024 it will generate the same numbers

	// notice that each call to IntN generates a different number however
	// if the seed is the same the sequence will be the same, this is
	// first generated number is equal in the 2 generators
	// gen 1: 42 -> 59 -> 84
	// gen 2: 42 -> 59 -> 84
}
