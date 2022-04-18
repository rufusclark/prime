package count

import "github.com/rufusclark/prime/primality"

// TrialDivision is a function that calculates a given number of prime number from 0 and returns them in a slice
func TrialDivision(count int) (primes []int) {
	if count < 1 {
		return []int{}
	}

	primes = make([]int, count)
	for i, c := 2, 0; c < count; i++ {
		if primality.TrialDivision(i) {
			primes[c] = i
			c++
		}
	}
	return primes
}
