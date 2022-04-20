package domain

// TrialDivision is a function that calculates all the primes between start and end (inclusive) and returns them in a slice
func TrialDivision(start, end int) (primes []int) {
	primes = make([]int, 0)

	isPrime := func(n int) bool {
		if n < 2 {
			return false
		}

		for i := 2; i < n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}

	for i := start; i <= end; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}
