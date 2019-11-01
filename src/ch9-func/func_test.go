package ch9_func

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func getTwoRandInt() (int, int) {
	return rand.Intn(10), rand.Intn(5)
}

func TestMultReturns(t *testing.T) {
	a, b := getTwoRandInt()
	t.Log(a, b)
}

func timeSpent(inner func(num int) int) func(num int) int {
	return func(num int) int {
		start := time.Now()
		ret := inner(num)
		fmt.Printf("spent time %f seconds", time.Since(start).Seconds())
		return ret
	}
}

func TestTimeSpent(t *testing.T) {
	t.Log(timeSpent(func1)(2))
}

func func1(num int) int {
	time.Sleep(1000)
	return num
}
