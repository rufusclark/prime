package count

import (
	"github.com/rufusclark/prime/primality"
)

// TrailDivision is a function that calculates a given number of prime number from 0 and returns them in a slice
func TrailDivision(count int) (primes []int) {
	if count < 1 {
		return []int{}
	}

	primes = make([]int, count)
	for i, c := 2, 0; c < count; i++ {
		if primality.TrailDivision(i) {
			primes[c] = i
			c++
		}
	}
	return primes
}
