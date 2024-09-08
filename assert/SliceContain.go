package assert

import (
	"slices"
	"testing"
)

func SliceContain[T comparable](t *testing.T, slice []T, x T) {
	if !slices.Contains(slice, x) {
		t.Error(slice, "does not contain", x)
	}
}
