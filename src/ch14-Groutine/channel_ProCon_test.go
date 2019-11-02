package ch14_Groutine

import (
	"fmt"
	"sync"
	"testing"
)

func Produce(chn chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			chn <- i
			fmt.Println("produce data:", i)
		}
		close(chn) // channel生产消息了 记得正常关闭。。。。。。。。。
		wg.Done()
	}()
}

func Consume(chn chan int, wg *sync.WaitGroup, name string) {
	go func() {
		for i := 0; i < 7; i++ {
			if data, ok := <-chn; ok { //生产者执行了close(chn) 才能这样判断，否则无数据可消费的时候会报错：all goroutines are asleep - deadlock!
				fmt.Println(name, " consume data:", data)
			} else {
				fmt.Println(name, " break at i=", i)
				break
			}

		}
		wg.Done()
	}()
}

func TestChannel(t *testing.T) {
	chn := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	Produce(chn, &wg)
	wg.Add(1)
	Consume(chn, &wg, "cons1")
	wg.Add(1)
	Consume(chn, &wg, "cons2")
	wg.Wait()
}
