package primality

import (
	"fmt"
	"testing"
)

func TestTrailDivision(t *testing.T) {
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
		testname := fmt.Sprintf("github.com/rufusclark/prime/primality/TrailDivision(%d)", test.n)

		t.Run(testname, func(t *testing.T) {
			got := TrailDivision(test.n)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
