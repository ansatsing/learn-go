package main

import (
	"fmt"
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

func main() {
	go spinner(10 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibnacci(%d)=%d\n", n, fibN)
}
