package tool

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"testing"
)

// createRandomIntSlice is a function that returns a slice of length < n with random ints
func createRandomIntSlice(n int) (ints []int) {
	n = rand.Intn(n) + 1
	ints = make([]int, n)
	last := 0
	for i := 0; i < n; i++ {
		ints[i] = last + rand.Intn(64)
		last = ints[i]
		// ints = append(ints, rand.Intn(math.MaxInt32))
	}
	return ints
}

// createRandomByteSlice is a function that returns a slice of length < n with random bytes
func createRandomByteSlice(n int) (bytes []byte) {
	n = rand.Intn(n)
	for i := 0; i < n+1; i++ {
		bytes = append(bytes, byte(rand.Intn(math.MaxInt8)))
	}
	return
}

func TestMarshal(t *testing.T) {
	rand.Seed(0)

	for i := 0; i < 64; i++ {
		want := createRandomIntSlice(8)
		testname := fmt.Sprintf("%v = unmarshal(marshal(%v))", want, want)

		t.Run(testname, func(t *testing.T) {
			got := unmarshal(marshal(want))
			if !reflect.DeepEqual(want, got) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	rand.Seed(0)

	for i := 0; i < 64; i++ {
		want := createRandomIntSlice(8)
		testname := fmt.Sprintf("%v = deecode(encode(%v))", want, want)

		t.Run(testname, func(t *testing.T) {
			got := decode(encode(want))
			if !reflect.DeepEqual(want, got) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func TestCompress(t *testing.T) {
	rand.Seed(0)

	for i := 0; i < 64; i++ {
		want := createRandomByteSlice(8)
		testname := fmt.Sprintf("%v = decompress(compress(%v))", want, want)

		t.Run(testname, func(t *testing.T) {
			compressed, _ := compress(want)
			got, _ := decompress(compressed)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func TestMarshalEncodeCompress(t *testing.T) {
	rand.Seed(0)

	for i := 0; i < 64; i++ {
		want := createRandomIntSlice(8)
		testname := fmt.Sprintf("%v = unmarshal(decode(decompress(compress(encode(marshal(%v))))))", want, want)

		t.Run(testname, func(t *testing.T) {
			masrshalled := marshal(want)
			encoded := encode(masrshalled)
			compressed, _ := compress(encoded)
			decompressed, _ := decompress(compressed)
			decoded := decode(decompressed)
			got := unmarshal(decoded)

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func BenchmarkMarshalEncodeCompress(b *testing.B) {
	rand.Seed(0)

	for i := 0; i < 8; i++ {
		want := createRandomIntSlice(8)
		testname := fmt.Sprintf("%v = prime > compreesed bytes > prime", want)

		b.Run(testname, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				masrshalled := marshal(want)
				encoded := encode(masrshalled)
				compressed, _ := compress(encoded)
				decompressed, _ := decompress(compressed)
				decoded := decode(decompressed)
				got := unmarshal(decoded)
				_ = got
			}
		})
	}
}
