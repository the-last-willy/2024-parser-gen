package assert

import "testing"

func Nil[T any](t *testing.T, a *T) {
	if a != nil {
		t.Error("Must be nil")
	}
}
