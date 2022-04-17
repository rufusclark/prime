package domain

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrimeFactor(t *testing.T) {
	tests := []struct {
		end  int
		want []int
	}{
		{0, []int{}},
		{-1, []int{}},
		{2, []int{2}},
		{3, []int{2, 3}},
		{5, []int{2, 3, 5}},
		{10, []int{2, 3, 5, 7}},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("github.com/rufusclark/prime/generator/domain/PrimeFactor(%d)", test.end)

		t.Run(testname, func(t *testing.T) {
			got := PrimeFactor(test.end)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
