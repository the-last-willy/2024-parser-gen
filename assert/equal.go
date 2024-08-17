package assert

import "testing"

func Equal[T comparable](t *testing.T, a, b T) {
	if a != b {
		t.Error("Got", a, ", Expected", b)
	}
}
