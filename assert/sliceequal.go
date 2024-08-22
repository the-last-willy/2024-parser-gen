package assert

import (
	"fmt"
	"testing"
)

func SliceEqual[T comparable](t *testing.T, a, b []T) {
	if len(a) != len(b) {
		t.Errorf("Slices have different lengths: %d vs %d", len(a), len(b))
	}
	for i := range a {
		if a[i] != b[i] {
			t.Errorf("Slices have different elements at index %d: %s vs %s", i, fmt.Sprint(a[i]), fmt.Sprint(b[i]))
		}
	}
}
