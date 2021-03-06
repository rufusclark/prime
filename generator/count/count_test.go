package count

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

// TestFuzzCompareCount is function that compares the results from different generators based on count
func TestFuzzCompareCount(t *testing.T) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		n := rand.Intn(1000)
		testname := fmt.Sprintf("CompareCount PrimeFactor(%d)==TrialDivision(%d)", n, n)

		t.Run(testname, func(t *testing.T) {
			if !reflect.DeepEqual(PrimeFactor(n), TrialDivision(n)) {
				t.Errorf("PrimeFactor(%d) != TrialDivision(%d)", n, n)
			}
		})
	}
}
