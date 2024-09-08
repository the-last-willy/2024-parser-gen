package assert

import "testing"

func True(t *testing.T, a bool) {
	if !a {
		t.Error("not true")
	}
}
