package ch14_Groutine

import (
	"fmt"
	"testing"
	"time"
)

func spinner(tim time.Duration) {
	for {
		for _, v := range `-\|/` {
			fmt.Printf("\r%c", v)
			time.Sleep(tim)
		}
	}

}
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func TestSpinner(t *testing.T) {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibnacci(%d)=%d\n", n, fibN)
}
