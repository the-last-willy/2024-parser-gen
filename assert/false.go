package assert

import "testing"

func False(t *testing.T, a bool) {
	if a {
		t.Error("not false")
	}
}
