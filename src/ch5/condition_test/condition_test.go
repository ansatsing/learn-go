package condition_test

import "testing"

func TestIf(t *testing.T) {
	if a := 1 == 1; a {
		t.Log(a)
	}
}
