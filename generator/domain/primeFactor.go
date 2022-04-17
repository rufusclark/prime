package domain

import "math"

// PrimeFactor is a function that calculates all prime number between 0 and end (inclusive) by checking if successive prime numbers are factors and returns them in a slice
func PrimeFactor(end int) (primes []int) {
	if end < 2 {
		return []int{}
	}

	primes = make([]int, 1)
	primes[0] = 2

	// utility function to check each number against previous primes
	isPrime := func(n int) bool {
		limit := int(math.Sqrt(float64(n)))

		for _, v := range primes {
			if n%v == 0 {
				return false
			}

			if v > limit {
				return true
			}
		}
		return true
	}

	for i := 3; i <= end; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	return primes
}
