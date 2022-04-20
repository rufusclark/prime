package tool

import (
	"bytes"
	"compress/gzip"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

const maxBytes = 8 // maximum number of bytes per number/int

// compress is a function that compress a byte slice using gzip and returns the compressed byte slice
func compress(data []byte) (compressed []byte, err error) {

	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	_, err = zw.Write(data)
	if err != nil {
		log.Fatal(err)
	}

	if err := zw.Close(); err != nil {
		log.Fatal(err)
	}

	return buf.Bytes(), err
}

// decompress is a function that decompress a byte slice using gzip and returns the uncompressed byte slice
func decompress(compressed []byte) (data []byte, err error) {

	buf := bytes.NewBuffer(compressed)

	zr, err := gzip.NewReader(buf)
	if err != nil {
		log.Fatal(err)
	}

	data, err = ioutil.ReadAll(zr)

	if err := zr.Close(); err != nil {
		log.Fatal(err)
	}

	return
}

// writeBytes is a function that writes bytes to local file (filename).
// this overwrites the file
func writeBytes(filename string, bytes []byte) (n int, err error) {

	f, err := os.OpenFile(filename, os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	n, err = f.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
	return n, err
}

// readBytes is a function that reads the contents of local file (filename) and returns them
func readBytes(filename string) (bytes []byte, err error) {

	bytes, err = os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return bytes, err
}

// marshal is a function that converts a slice of primes into a slice of differences of primes.
// ints must be in increasing order - this is assumed as no checks are made for this
func marshal(primes []int) (mPrimes []int) {

	mPrimes = make([]int, len(primes))
	low := 0

	for i := 0; i < len(primes); i++ {
		mPrimes[i] = primes[i] - low
		low = primes[i]
	}

	return mPrimes
}

// unmarshal is a function that converts a slice of differences of primes to a slice of primes
func unmarshal(mPrimes []int) (primes []int) {

	if l := len(mPrimes); l == 0 {
		return
	}

	primes = make([]int, len(mPrimes))
	primes[0] = mPrimes[0]

	for i := 1; i < len(primes); i++ {
		primes[i] = primes[i-1] + mPrimes[i]
	}

	return primes
}

// encodeOnce is a function that encodes a single int value to a byte slice
func encodeOnce(value int) (bytes []byte, err error) {

	if value < 0 {
		return bytes, errors.New("encoding error: value must be positve")
	}

	if value > 127 {
		// add final byte (no flag)
		bytes = append([]byte{byte((value % 128))}, bytes...)
		value = value / 128
	} else {
		// return bytes if value can be represented by 7 bits
		bytes = append(bytes, byte(value))
		return bytes, nil
	}

	for i := 1; i <= maxBytes; i++ {
		if value > 127 {
			bytes = append([]byte{byte((value % 128) + 128)}, bytes...)
			value = value / 128
		} else {
			bytes = append([]byte{byte((value + 128))}, bytes...)
			return bytes, nil
		}
	}

	return bytes, errors.New("encoding error: unexpected behaviour: reached end of function")
}

// decodeOnce is a function that returns the decodes a byte slice and returns (n = number of bytes for value, value = decoded value, err = error)
func decodeOnce(bytes []byte) (n, value int, err error) {

	for i := 0; i < len(bytes); i++ {
		if bytes[i] > 127 {
			value = (value + int(bytes[i]) - 128) * 128
			// check if number is larger then possible under protocol implementation
			if i >= maxBytes-1 {
				return 0, 0, errors.New("decoding failed: max implemented number is 7.2057594e+16")
			}
		} else {
			value += int(bytes[i])
			n = i + 1
			return n, value, nil
		}
	}
	return 0, 0, errors.New("decoding failed: end of byte slice - no encoded value")
}

func encode(values []int) (bytes []byte) {
	for _, value := range values {
		encoded, err := encodeOnce(value)
		if err != nil {
			log.Fatal(err)
		}
		bytes = append(bytes, encoded...)
	}

	return bytes
}

func decode(bytes []byte) (values []int) {

	max := len(bytes)
	high, span := 0, 0

	for low := 0; low < max; low++ {

		span = max - low
		if span > maxBytes {
			high = low + maxBytes
		} else {
			high = low + span
		}

		n, value, err := decodeOnce(bytes[low:high])
		if err != nil {
			log.Fatal(err)
		}

		low += n - 1
		values = append(values, value)
	}

	return values
}
