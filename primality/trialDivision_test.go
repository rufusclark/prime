package primality

import (
	"fmt"
	"testing"
)

func BenchmarkTrialDivision(b *testing.B) {
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
		benchname := fmt.Sprintf("TrialDivision(%d)", test.n)

		b.Run(benchname, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				TrialDivision(test.n)
			}
		})
	}
}

func TestTrialDivision(t *testing.T) {
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
		testname := fmt.Sprintf("github.com/rufusclark/prime/primality/TrialDivision(%d)", test.n)

		t.Run(testname, func(t *testing.T) {
			got := TrialDivision(test.n)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
