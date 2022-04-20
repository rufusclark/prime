package primality

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

// TestFuzzCompareDomain is function that compares the results from different generators based on count
func TestFuzzComparePrimality(t *testing.T) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		n := rand.Intn(100000)
		testname := fmt.Sprintf("ComparePrimality TrialDivision(%d) == HybridCache(%d)", n, n)

		t.Run(testname, func(t *testing.T) {
			got1 := TrialDivision(n)
			got2 := HybridCache(n)
			if !reflect.DeepEqual(got1, got2) {
				t.Errorf("TrialDivision(%d) != HybridCache(%d)", n, n)
			}
		})
	}
}
