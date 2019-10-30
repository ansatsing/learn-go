package loop_test

import "testing"

func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
}

func TestForLoop(t *testing.T) {
	n := 9
	for i := 0; i < n; i++ {
		t.Log(i)
	}
}

func TestDeadLoop(t *testing.T) {
	i := 0
	for {
		t.Log(i)
		i++
	}
}
