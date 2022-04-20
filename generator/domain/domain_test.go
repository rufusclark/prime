package domain

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

// TestFuzzCompareDomain is function that compares the results from different generators based on count
func TestFuzzCompareDomain(t *testing.T) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		n := rand.Intn(10000)
		testname := fmt.Sprintf("CompareCount PrimeFactor(%d)==TrialDivision(%d)==Seo(%d)", n, n, n)

		t.Run(testname, func(t *testing.T) {
			got1 := PrimeFactor(n)
			got2 := Seo(n)
			got3 := TrialDivision(0, n)
			if !(reflect.DeepEqual(got1, got2) && reflect.DeepEqual(got1, got3)) {
				t.Errorf("PrimeFactor(%d) != TrialDivision(%d) != Seo(%d)", n, n, n)
			}
		})
	}
}
