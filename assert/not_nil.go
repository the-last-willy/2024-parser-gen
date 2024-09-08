package assert

import "testing"

func NotNil[T any](t *testing.T, a *T) {
	if a == nil {
		t.Error("unexpected nil")
	}
}
