package domain

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTrialDivision(t *testing.T) {
	tests := []struct {
		start, end int
		want       []int
	}{
		{0, 0, []int{}},
		{0, 2, []int{2}},
		{0, 5, []int{2, 3, 5}},
		{2, 0, []int{}},
		{-10, 2, []int{2}},
		{10, -2, []int{}},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("github.com/rufusclark/prime/generator/domain/TrialDivision(%d, %d)", test.start, test.end)

		t.Run(testname, func(t *testing.T) {
			got := TrialDivision(test.start, test.end)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
