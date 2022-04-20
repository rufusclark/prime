package tool

import (
	"errors"
	"log"
	"os"

	"github.com/rufusclark/prime/generator/domain"
)

// These functions are not performance optomized and are deliberately broken down into discrete function

// WritePrimes is a function that writes primes to a file
//	params
// 		filename = filename relative to current execution
// 		primes = list of primes in increasing order - this is assumed and not checked, ignoring this will cause unexpected behabiour and errors
//  returns
// 		n = number of bytes written,
// 		err = error or nil
func WritePrimes(filename string, primes []int) (n int, err error) {

	marshalledPrimes := marshal(primes)

	encodedPrimes := encode(marshalledPrimes)

	compressedPrimes, err := compress(encodedPrimes)
	if err != nil {
		log.Fatal(err)
	}

	n, err = writeBytes(filename, compressedPrimes)
	if err != nil {
		log.Fatal(err)
	}

	return
}

// ReadPrimes is a function that reads primes from a file
func ReadPrimes(filename string) (primes []int, err error) {

	compressedPrimes, err := readBytes(filename)
	if err != nil {
		log.Fatal(err)
	}

	encodedPrimes, err := decompress(compressedPrimes)
	if err != nil {
		log.Fatal(err)
	}

	masrhalledPrimes := decode(encodedPrimes)

	primes = unmarshal(masrhalledPrimes)

	return
}

// LoadCache is a function which will load the prime cache or create a new one if it doesn't exist
// (see GeneratePrimeCache for more details)
func LoadCache(filename string, exponent int) (primes []int, err error) {
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		primes, err := GenerateCache(filename, exponent)
		return primes, err
	} else {
		return ReadPrimes(filename)
	}
}

// GeneratePrimeCache generates a prime cache and saved it in filename
// params
// 			filename = string of filename
//			exponent = int representing exponent of 2 primes to generate (ie. 2^n = generated count)
// returns
//			err = error is present else nil
func GenerateCache(filename string, exponent int) (primes []int, err error) {
	if exponent < 1 || exponent > 22 {
		return primes, errors.New("invalid exponent: must be between 1 and 22 (inclusive)")
	}

	// convert exponent of 2 into the domain for the generator.
	// these values are precomputed and allow an efficient sieve to be implemented for the generation of these primes.
	// this results in the cache size being optimal for a binary search where exponent is a proxy for passes
	exponent2domain := []int{0, 3, 7, 19, 53, 131, 311, 719, 1619, 3671, 8161, 17863, 38873, 84017, 180503, 386093, 821641, 1742537, 3681131, 7754077, 16290047, 34136029, 71378569}

	n := exponent2domain[exponent]

	primes = domain.Seo(n)
	_, err = WritePrimes(filename, primes)

	return primes, err
}
