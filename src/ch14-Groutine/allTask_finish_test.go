package ch14_Groutine

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"
	"time"
)

//利用csp实现类似WaitGroup的功能

func doTaskReturnRandData(i int64) int {
	j := rand.New(rand.NewSource(time.Now().UnixNano() + i)).Intn(100)
	fmt.Println("task ", i, " generate data:", j)
	return j
}

func getAllTaskData(i int64) []int {
	chn := make(chan int, i)
	var j int64
	for j = 0; j < i; j++ {
		go func(e int64, ch chan int) {
			ch <- doTaskReturnRandData(e)
		}(j, chn)
	}
	time.Sleep(time.Millisecond * 100)
	ret := []int{}
	for j = 0; j < i; j++ {
		ret = append(ret, <-chn)
	}
	return ret
}

func TestAllTaskFinish(t *testing.T) {
	t.Log("before->current goroutines size:", runtime.NumGoroutine())
	t.Log(getAllTaskData(10))
	time.Sleep(time.Second * 1)
	t.Log("after->current goroutines size:", runtime.NumGoroutine())
}
