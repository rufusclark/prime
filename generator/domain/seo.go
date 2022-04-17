package domain

import "math"

// Seo (sieve of eratosthenes) is a function that returns a slice with all primes in the range 0 to end (inclusive)
func Seo(end int) (primes []int) {
	if end < 2 {
		return []int{}
	}

	sieve := make([]bool, end+1)
	count := 0

	for i := 2; i <= int(math.Sqrt(float64(end))); i++ {
		if !sieve[i] {
			count++
		}
		for j := 2 * i; j <= end; j += i {
			sieve[j] = true
		}
	}

	primes = make([]int, 0)
	for i, v := range sieve {
		if !v {
			primes = append(primes, i)
		}
	}
	primes = primes[2:]

	return primes
}
