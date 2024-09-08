package assert

import "testing"

func Implements[I any](t *testing.T, v any) {
	_, ok := v.(I)
	if !ok {
		t.Error("value does not implement interface")
	}
}
