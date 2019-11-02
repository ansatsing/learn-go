package ch14_Groutine

import (
	"sync"
	"testing"
	"time"
)

//共享内存并发机制
//Mutex ==== synchonized
//WaitGroup  ==== join

func TestThreadUnsafe(t *testing.T) {
	counter := 0
	for i := 0; i < 200; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Second * 2)
	t.Logf("counter:%d", counter) //期望结果是200，几乎每次结果都不一样
}

func TestThreadSafeWithMutex(t *testing.T) {
	//mut :=sync.Mutex  //不能这样
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 200; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(time.Second * 2)   //这样 是不对的。。
	t.Logf("counter:%d", counter) //期望结果是200，几乎每次结果都不一样
}

func TestThreadSafeWithMutexAndWaitGroup(t *testing.T) {
	//mut :=sync.Mutex  //不能这样
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("counter:%d", counter) //期望结果是200，几乎每次结果都不一样
}
