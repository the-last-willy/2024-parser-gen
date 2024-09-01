package assert

import "testing"

func Panic(t *testing.T, f func()) {
	defer func() {
		if recover() == nil {
			t.Error("assert.Panic: did not panic")
		}
	}()

	f()
}
