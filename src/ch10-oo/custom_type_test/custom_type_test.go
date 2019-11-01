package custom_type_test

import (
	"fmt"
	"testing"
	"time"
)

type MyFunc func(num int) int

func TimeSpent(inner MyFunc) MyFunc {
	return func(num int) int {
		start := time.Now()
		ret := inner(num)
		fmt.Println("time spent：", time.Since(start).Seconds())
		return ret
	}

}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	tsSF := TimeSpent(slowFun)
	t.Log(tsSF(10))
}
