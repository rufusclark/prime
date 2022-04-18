package primality

import (
	"math"
)

// binarySearch is a function which returns a bool representing whether or not a given int is in the global []int slice cache
func binarySearch(n int) bool {
	low, high, mid := 0, len(cache)-1, 0
	for low <= high {
		mid = (low + high) / 2
		if cache[mid] < n {
			low = mid + 1
		} else if cache[mid] > n {
			high = mid - 1
		} else {
			return true
		}
	}
	return false
}

// HybridCache is a function which returns a bool representing whether or not a value is prime
// This works by first checking if it is in a prime, then checking the cached list as factors then using trial division if not previous methods are conclusive
// On first execution this creates the cache/loads it into memory
func HybridCache(n int) bool {
	// check if less than 2
	if n < 2 {
		return false
	}

	// check if n is in cache
	max := cache[len(cache)-1]
	if n < max {
		return binarySearch(n)
	}

	// check all posible prime factors
	sqrtn := int(math.Sqrt(float64(n)))
	if sqrtn < max {
		for _, v := range cache {
			if n%v == 0 {
				return false
			}
		}
		return true
	}

	// check all possible factors then use an exhaustive test
	for _, v := range cache {
		if n%v == 0 {
			return false
		}
	}
	for i := sqrtn; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
