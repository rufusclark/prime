package domain

import "math"

// Seo (sieve of eratosthenes) is a function that returns a slice with all primes in the range 0 to end (inclusive)
func Seo(end int) (primes []int) {
	if end < 2 {
		return []int{}
	}

	sieve := make([]bool, end+1)          // sieve slice
	primes = make([]int, 0)               // output slice with primes
	limit := int(math.Sqrt(float64(end))) // last prime factor to sieve

	for i := 2; i <= limit; i++ {
		// checking if element in sieve is false and hence prime
		if !sieve[i] {
			primes = append(primes, i)
			for j := 2 * i; j <= end; j += i {
				sieve[j] = true
			}
		}
	}

	// add the remaining non multiples to primes slice
	for i := limit; i <= end; i++ {
		if !sieve[i] {
			primes = append(primes, i)
		}
	}

	return primes
}
