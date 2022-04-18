package domain

import "github.com/rufusclark/prime/primality"

// TrialDivision is a function that calculates all the primes between start and end (inclusive) and returns them in a slice
func TrialDivision(start, end int) (primes []int) {
	primes = make([]int, 0)

	for i := start; i <= end; i++ {
		if primality.TrialDivision(i) {
			primes = append(primes, i)
		}
	}
	return primes
}
