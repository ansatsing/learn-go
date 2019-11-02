package ch14_Groutine

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

//多路选择和超时

func service() int {
	time.Sleep(time.Millisecond * 600)
	return rand.Int()
}

func future() chan int {
	chn := make(chan int, 1)
	go func() {
		chn <- service()
	}()
	return chn
}

func TestSelectTimeout(t *testing.T) {
	select {
	case data := <-future():
		t.Log("data:", data)
	case <-time.After(time.Millisecond * 800):
		panic("TimeOut:")
	}
}

/////////////////// ////////////////////////
//
//
// 总共3个任务，任务1和任务2耗时，任务3不耗时，但任务1的结果必须传给任务3,
// 任务2根其他2个任务毛关系,任务1要做超时控制，超时了就不执行任务3

func doTask1() int {
	time.Sleep(time.Millisecond * 600)
	return rand.Intn(100)
}

func doTask2() {
	time.Sleep(time.Millisecond * 680)
	fmt.Println("task2 has finished")
}

func doTask3(i int) {
	vle := 10 + i
	fmt.Println("task3 has finished,data:", vle)
}

func doTask1WithFuture() chan int {
	chn := make(chan int, 1)
	go func() {
		chn <- doTask1()
	}()
	return chn
}

func doTask1WithTimeOut(tim time.Duration) (int, error) {
	dur := time.Millisecond * tim
	select {
	case data := <-doTask1WithFuture():
		return data, nil
	case <-time.After(dur):
		return 0, errors.New("Time out.")
	}
}

func TestCLient(t *testing.T) {
	data, err := doTask1WithTimeOut(700)
	doTask2()
	if err == nil {
		doTask3(data)
	} else {
		fmt.Println("task2 time out")
	}

}
