package ch14_Groutine

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"
	"time"
)

//就获取最先完成的那个任务的数据。。。。

func doTask() int {
	//i:=rand.Intn(10)
	i := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)
	var dur time.Duration
	dur = time.Millisecond * time.Duration(i)
	time.Sleep(dur)
	return i
}

func TestFirst(t *testing.T) {
	num := 6
	chn := make(chan int, num)
	for i := 0; i < num; i++ {
		go func(j int, ch chan int) {
			ret := doTask()
			fmt.Println(j, " data:", ret)
			ch <- ret
		}(i, chn)
	}

	t.Log("result:", <-chn)
	time.Sleep(time.Second * 2)
}

///////////////
func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}

func FirstResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func TestFirstResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())

}
