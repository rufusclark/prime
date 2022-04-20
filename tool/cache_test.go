package tool

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReadWrite(t *testing.T) {
	const filename = "test.cache"
	want := make([]int, 0)

	for i := 0; i < 512; i += 64 {
		testname := fmt.Sprintf("ReadPrimes() = WritePrimes() %d", i)

		want = append(want, i)

		t.Run(testname, func(t *testing.T) {
			WritePrimes(filename, want)
			got, _ := ReadPrimes(filename)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
