package count

import "math"

// PrimeFactor is a function that calculates a given number of prime number from 0 by checking if successive prime numbers are factors and returns them in a slice
func PrimeFactor(count int) (primes []int) {
	if count < 1 {
		return []int{}
	}

	primes = make([]int, count)
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

	for i, c := 3, 1; c < count; i++ {
		if isPrime(i) {
			primes[c] = i
			c++
		}
	}

	return primes
}
