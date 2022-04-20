package count

import (
	"fmt"
	"reflect"
	"testing"
)

func BenchmarkTrialDivision(b *testing.B) {
	inputs := []int{10, 100, 1000, 10000}

	for _, input := range inputs {
		benchname := fmt.Sprintf("TrialDivision(%d)", input)

		b.Run(benchname, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				TrialDivision(input)
			}
		})
	}
}

func TestTrialDivision(t *testing.T) {
	tests := []struct {
		count int
		want  []int
	}{
		{0, []int{}},
		{-10, []int{}},
		{1, []int{2}},
		{2, []int{2, 3}},
		{5, []int{2, 3, 5, 7, 11}},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("TrialDivision(%d)", test.count)

		t.Run(testname, func(t *testing.T) {
			got := TrialDivision(test.count)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
