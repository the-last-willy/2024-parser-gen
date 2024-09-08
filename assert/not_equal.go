package assert

import "testing"

func NotEqual[T comparable](t *testing.T, a, b T) {
	if a == b {
		t.Error("Not equal failed =,", a)
	}
}
