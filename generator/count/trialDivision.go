package count

// TrialDivision is a function that calculates a given number of prime number from 0 and returns them in a slice
func TrialDivision(count int) (primes []int) {
	if count < 1 {
		return []int{}
	}

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

	primes = make([]int, count)
	for i, c := 2, 0; c < count; i++ {
		if isPrime(i) {
			primes[c] = i
			c++
		}
	}
	return primes
}
