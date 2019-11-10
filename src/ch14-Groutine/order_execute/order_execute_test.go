package order_execute

import (
	"fmt"
	"os"
	"sync/atomic"
	"testing"
	"time"
)

//顺序执行：0 1 2 3 4 5 6 7 8 9
func TestOrderBYChannel(t *testing.T) {
	chn := make(chan struct{}, 1)
	for i := 0; i < 10; i++ {
		chn <- struct{}{}
		go func(i int) {
			s := <-chn
			fmt.Println(i, s)
		}(i)
	}
	time.Sleep(time.Second * 2)
}

//顺序执行：0 1 2 3 4 5 6 7 8 9
func TestOrderBYChannel1(t *testing.T) {
	chn := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
			chn <- struct{}{}
		}(i)
		<-chn
	}
	time.Sleep(time.Second * 2)
	os.PathError{}
}

//利用自旋spinning
func TestOrderBYAtomic(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fn := func() { fmt.Println(i) }
			trigger(i, fn)
		}(i)
	}
	trigger(10, func() {})
}

var count = int32(0)

func trigger(i int, fn func()) {
	for {
		if atomic.LoadInt32(&count) == int32(i) {
			fn()
			atomic.AddInt32(&count, 1)
			break
		}
		time.Sleep(time.Microsecond * 100) //没这行代码直接ide卡死
	}
}
