package primality

import (
	"fmt"
	"testing"
)

func BenchmarkHybridCache(b *testing.B) {
	tests := []struct {
		n    int
		want bool
	}{
		{0, false},
		{-10, false},
		{2, true},
		{11, true},
		{97, true},
		{100, false},
		{10000, false},
	}

	for _, test := range tests {
		benchname := fmt.Sprintf("HybridCache(%d)", test.n)

		b.Run(benchname, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HybridCache(test.n)
			}
		})
	}
}

func TestHybridCache(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{0, false},
		{-10, false},
		{2, true},
		{11, true},
		{97, true},
		{100, false},
		{10000, false},
		{100000, false},
		{1000003, true},
		{15485863, true},
		{999999937, true},
		{999999938, false},
		{2147483647, true},
		{2147483648, false},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("HydridCache(%d)", test.n)

		t.Run(testname, func(t *testing.T) {
			got := HybridCache(test.n)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
