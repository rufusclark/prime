package domain

import (
	"fmt"
	"reflect"
	"testing"
)

func BenchmarkSeo(b *testing.B) {
	inputs := []int{10, 100, 1000, 10000}

	for _, input := range inputs {
		benchname := fmt.Sprintf("Seo(%d)", input)

		b.Run(benchname, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Seo(input)
			}
		})
	}
}

func TestSeo(t *testing.T) {
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
		testname := fmt.Sprintf("Seo(%d)", test.end)

		t.Run(testname, func(t *testing.T) {
			got := PrimeFactor(test.end)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
